package core

import (
	"encoding/json"
	"errors"
)

type Transaction struct {
	IsCoinbase bool            `json:"isCoinbase"`
	Data       TransactionData `json:"data"`
	Signature  Signature       `json:"signature"`
}

type TransactionData struct {
	Sender   PublicKey `json:"publicKey"`
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

func (t Transaction) VerifySignature() bool {
	data, _ := t.MarshalData()
	return t.Signature.Verify(t.Data.Sender, data)
}

func (t Transaction) IsFieldEmpty() bool {
	return t.Data.Sender == EmptyPublicKey || t.Data.Receiver == EmptyPublicKey || t.Data.Amount == 0 || t.Signature == EmptySignature
}

func (t Transaction) MarshalData() ([]byte, error) {
	if t.IsFieldEmpty() {
		return []byte{}, errors.New("fields missing")
	}

	j, err := json.Marshal(t.Data)
	if err != nil {
		return []byte{}, err
	}

	return j, nil
}

func (t Transaction) Marshal() ([]byte, error) {
	if t.IsFieldEmpty() {
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
	if t.IsFieldEmpty() {
		return Transaction{}, errors.New("fields missing")
	}

	return t, nil
}
