package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/chrigeeel/matura-blockchain/core"
)

type Config struct {
	PrivateKey core.PrivateKey `json:"privateKey"`
}

var (
	config Config
)

func init() {
	f, err := ioutil.ReadFile("../config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}
}

const RPCURL = ""

func main() {

	privateKey := config.PrivateKey

	receiver := core.MustPrivateKeyFromBase58("5RargVQn8LcE7PvuE1wZnzvZS63ZYMVMxgPBbhatvVqHKUbeJCjLR8D4NgEqLr6NMrHcXeEyp44QwqT6f9sJc4AZ")

	//rpcClient := rpc.NewClient("http://3.66.6.144")

	fmt.Println(receiver.PublicKey())

	tx := core.NewTransaction(
		core.TransactionData{
			Sender:   privateKey.PublicKey(),
			Receiver: receiver.PublicKey(),
			Amount:   core.FractionsPerCoin * 1,
		},
	)

	signature, err := tx.Sign(privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(signature)

	fmt.Println(tx.DataToSign())
}
