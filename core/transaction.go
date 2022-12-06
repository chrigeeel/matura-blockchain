package core

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"math"
	"math/big"
)

type RPCTransaction struct {
	Status        string      `json:"status"`
	Confirmations int         `json:"confirmations"`
	Data          Transaction `json:"data"`
}

type Transaction struct {
	IsCoinbase bool            `json:"isCoinbase"`
	Data       TransactionData `json:"data"`
	Signature  Signature       `json:"signature"`
}

type TransactionData struct {
	Sender   PublicKey `json:"sender"`
	Receiver PublicKey `json:"receiver"`
	Amount   uint64    `json:"amount"`
	Nonce    uint64    `json:"nonce"`
}

func (t *Transaction) Sign(p PrivateKey) (Signature, error) {
	data, err := t.DataToSign()
	if err != nil {
		return Signature{}, err
	}

	signature, err := p.Sign(data[:])
	if err != nil {
		return Signature{}, err
	}

	t.Signature = signature
	return signature, nil
}

func (t Transaction) DataToSign() (Hash, error) {
	data, err := t.MarshalData()
	if err != nil {
		return Hash{}, err
	}

	return HashBytes(data), nil
}

func (t Transaction) VerifyTotal(regardMempool bool) bool {
	if !t.VerifySignature() {
		return false
	}

	if !t.VerifyNonce() {
		return false
	}

	if t.IsFieldEmpty() {
		return false
	}

	balance := ThisChain.GetBalance(t.Data.Sender, regardMempool)
	return t.Data.Amount <= balance
}

func (t Transaction) VerifySignature() bool {
	data, _ := t.DataToSign()
	return t.Signature.Verify(t.Data.Sender, data[:])
}

func (t Transaction) VerifyNonce() bool {
	for _, block := range ThisChain {
		for _, tx := range block.Transactions {
			if tx.Data.Nonce == t.Data.Nonce {
				return false
			}
		}
	}

	return true
}

func (t Transaction) IsFieldEmpty() bool {
	return t.Data.Sender == EmptyPublicKey || t.Data.Receiver == EmptyPublicKey || t.Data.Amount == 0 || t.Signature == EmptySignature
}

func (t Transaction) IsReadyToSign() bool {
	return t.Data.Sender == EmptyPublicKey || t.Data.Receiver == EmptyPublicKey || t.Data.Amount == 0
}

func (t Transaction) MarshalData() ([]byte, error) {
	if t.IsReadyToSign() {
		return []byte{}, errors.New("fields missing")
	}

	j, err := json.Marshal(t.Data)
	if err != nil {
		return []byte{}, err
	}

	return j, nil
}

func (t Transaction) Marshal() ([]byte, error) {
	if t.IsReadyToSign() {
		return []byte{}, errors.New("fields missing")
	}

	j, err := json.Marshal(t)
	if err != nil {
		return []byte{}, err
	}

	return j, nil
}

func UnmarshalTransaction(b []byte) (Transaction, error) {
	var t Transaction
	err := json.Unmarshal(b, &t)
	if err != nil {
		return Transaction{}, err
	}
	if t.IsReadyToSign() {
		return Transaction{}, errors.New("fields missing")
	}

	return t, nil
}

func NewTransaction(opts TransactionData) Transaction {
	opts.Nonce, _ = randint64()

	return Transaction{
		IsCoinbase: false,
		Data:       opts,
	}
}

func randint64() (uint64, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0, err
	}
	return val.Uint64(), nil
}
