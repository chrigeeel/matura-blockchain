package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	NodePort = 6923
)

var (
	PermanentPeers = []string{"3.122.2.210"}
)

func StartNodeServer() {
	go NodeServer()
}

func NodeServer() {
	app := fiber.New()

	app.Post("/block", func(c *fiber.Ctx) error {
		block := new(Block)

		if err := c.BodyParser(block); err != nil {
			return err
		}

		fmt.Println("new block")
		fmt.Println(block)

		err := AddBlock(*block, false)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/transaction", func(c *fiber.Ctx) error {
		transaction := new(Transaction)

		if err := c.BodyParser(transaction); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		fmt.Println("new transaction from broadcast")

		err := AddTransactionToMempool(*transaction)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/peer/+", func(c *fiber.Ctx) error {
		peer := c.Params("+")
		if peer == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		fmt.Printf("new peer %s", peer)
		WritePeer(peer)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/mempool", func(c *fiber.Ctx) error {
		return c.JSON(Mempool)
	})

	app.Get("/chain", func(c *fiber.Ctx) error {
		return c.JSON(ThisChain)
	})

	app.Get("/peers", func(c *fiber.Ctx) error {
		peers, err := ReadPeers()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(peers)
	})

	app.Listen(fmt.Sprintf(":%d", NodePort))
}

func BroadcastTransaction(transaction Transaction) error {
	err := AddTransactionToMempool(transaction)
	if err != nil {
		return err
	}

	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	self := GetSelfPeer()
	for _, peer := range peers {
		if peer == self {
			continue
		}
		go func(peer string) {
			url := fmt.Sprintf("http://%s:%d/transaction", peer, NodePort)
			j, err := json.Marshal(transaction)
			if err != nil {
				return
			}
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
			if err != nil {
				return
			}
			req.Header.Set("content-type", "application/json")
			http.DefaultClient.Do(req)
		}(peer)
		time.Sleep(time.Millisecond * 10)
	}

	return nil
}

func BroadcastBlock(block Block) error {
	fmt.Println("broadcasting block")
	err := AddBlock(block, true)
	if err != nil {
		panic(err)
	}

	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	self := GetSelfPeer()
	for _, peer := range peers {
		if peer == self {
			continue
		}
		go func(peer string) {
			url := fmt.Sprintf("http://%s:%d/block", peer, NodePort)
			j, err := json.Marshal(block)
			if err != nil {
				return
			}
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
			if err != nil {
				return
			}
			req.Header.Set("content-type", "application/json")
			http.DefaultClient.Do(req)
		}(peer)
	}

	return nil
}
