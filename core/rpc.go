package core

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	RPCPort = 80
)

func StartRPCServer() {
	go RPCServer()
}

func RPCServer() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/transaction", func(c *fiber.Ctx) error {
		transaction := new(Transaction)

		if err := c.BodyParser(transaction); err != nil {
			return JSONError(c, fiber.StatusBadRequest, err)
		}

		if transaction.IsFieldEmpty() || !transaction.VerifySignature() || !transaction.VerifyNonce() {
			return JSONError(c, fiber.StatusBadRequest, errors.New("invalid transaction"))
		}

		err := BroadcastTransaction(*transaction)
		if err != nil {
			return JSONError(c, fiber.StatusBadRequest, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "mempool",
			"data":   transaction,
		})
	})

	app.Get("/transaction/+", func(c *fiber.Ctx) error {
		signature, err := SignatureFromBase58(c.Params("+"))
		if err != nil {
			return JSONError(c, fiber.StatusBadRequest, err)
		}

		for _, tx := range Mempool {
			if tx.Signature == signature {
				return c.Status(fiber.StatusOK).JSON(RPCTransaction{
					Status: "mempool",
					Data:   tx,
				})
			}
		}

		for i, block := range ThisChain {
			for _, tx := range block.Transactions {
				if tx.Signature == signature {
					return c.Status(fiber.StatusOK).JSON(RPCTransaction{
						Status:        "confirmed",
						Confirmations: i,
						Data:          tx,
					})
				}
			}
		}

		return JSONError(c, fiber.StatusNotFound, errors.New("not found"))
	})

	app.Get("/+/transactions", func(c *fiber.Ctx) error {
		publicKey, err := PublicKeyFromBase58(c.Params("+"))
		if err != nil {
			return JSONError(c, fiber.StatusBadRequest, err)
		}

		var transactions []interface{}

		for _, tx := range Mempool {
			if tx.Data.Sender.Equals(publicKey) || tx.Data.Receiver.Equals(publicKey) {
				transactions = append(transactions, tx)
			}
		}

		for _, block := range ThisChain {
			for _, tx := range block.Transactions {
				if tx.Data.Sender.Equals(publicKey) || tx.Data.Receiver.Equals(publicKey) {
					transactions = append(transactions, tx)
				}
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"transactions": transactions,
		})
	})

	app.Get("/balance/+", func(c *fiber.Ctx) error {
		publicKey, err := PublicKeyFromBase58(c.Params("+"))
		if err != nil {
			return JSONError(c, fiber.StatusBadRequest, err)
		}

		balance := ThisChain.GetBalance(publicKey, true)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"balance": balance,
		})
	})

	app.Listen(fmt.Sprintf(":%d", RPCPort))
}

func JSONError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"message": err.Error(),
	})
}
