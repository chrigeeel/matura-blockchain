package main

import (
	"encoding/json"
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
	f, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}

	core.StartCore()
}

func main() {
	core.Mine(config.PrivateKey)
}
