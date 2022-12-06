package core

import (
	"errors"
	"fmt"
)

var (
	Mempool []Transaction
)

func AddTransactionToMempool(transaction Transaction) error {
	if transaction.IsFieldEmpty() || !transaction.VerifySignature() || !transaction.VerifyNonce() {
		return errors.New("invalid transaction")
	}

	fmt.Println("new transaction in mempool")

	Mempool = append(Mempool, transaction)

	return nil
}

func ClearMempool() {
	fmt.Println("clearing mempool")
	Mempool = []Transaction{}
}
