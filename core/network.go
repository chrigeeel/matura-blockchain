package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	NodePort = 6923
)

var (
	PermanentPeers = []string{"3.122.2.210", "3.124.44.37", "3.66.6.144"}
)

func NodeServer() {
	app := fiber.New()

	app.Post("/block", func(c *fiber.Ctx) error {
		block := new(Block)

		if err := c.BodyParser(block); err != nil {
			return err
		}

		err := AddBlock(*block)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendStatus(fiber.StatusOK)
	})
	app.Post("/transaction", func(c *fiber.Ctx) error {
		transaction := new(Transaction)

		if err := c.BodyParser(transaction); err != nil {
			return err
		}

		AddTransactionToMempool(*transaction)

		return c.SendStatus(fiber.StatusOK)
	})
	app.Post("/peer", func(c *fiber.Ctx) error {
		peer := c.FormValue("peer")
		if peer == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

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
}

func BroadcastTransaction(transaction Transaction) error {
	AddTransactionToMempool(transaction)

	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(peers))
	for _, peer := range peers {
		go func(wg *sync.WaitGroup, peer string) {
			defer wg.Done()

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
		}(&wg, peer)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()

	return nil
}

func BroadcastBlock(block Block) error {
	err := AddBlock(block)
	if err != nil {
		panic(err)
	}

	fmt.Println("done adding")

	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	for _, peer := range peers {
		url := fmt.Sprintf("http://%s:%d/block", peer, NodePort)
		j, err := json.Marshal(block)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
		if err != nil {
			return err
		}
		req.Header.Set("content-type", "application/json")
		http.DefaultClient.Do(req)
	}

	return nil
}

func ReadPeers() ([]string, error) {
	var peers []string

	f, err := ioutil.ReadFile("./peers.json")
	if err != nil {
		return []string{}, err
	}

	err = json.Unmarshal(f, &peers)
	return peers, err
}

func WritePeer(peer string) error {
	peers, err := ReadPeers()
	if err != nil {
		return err
	}

	for _, p := range peers {
		if p == peer {
			return errors.New("peer already registered")
		}
	}

	peers = append(peers, peer)
	j, err := json.MarshalIndent(peers, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./peers.json", j, 0644)
}
