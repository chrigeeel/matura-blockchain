package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type Chain []Block

var (
	ThisChain Chain
	Mempool   []Transaction
)

func fetch(peer string) error {
	baseUrl := fmt.Sprintf("http://%s:%d", peer, NodePort)
	resp, err := http.Get(fmt.Sprintf("%s/chain", baseUrl))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &ThisChain)
	if err != nil {
		return err
	}

	err = WriteChain(ThisChain)
	if err != nil {
		return err
	}

	resp, err = http.Get(fmt.Sprintf("%s/mempool", baseUrl))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Mempool)
	if err != nil {
		return err
	}

	resp, err = http.Get(fmt.Sprintf("%s/peers", baseUrl))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var p []string
	err = json.Unmarshal(body, &p)
	if err != nil {
		return err
	}
	for _, pe := range p {
		WritePeer(pe)
	}

	return nil
}

func init() {
	peers, err := ReadPeers()
	if err != nil {
		panic(err)
	}
	fmt.Println(peers)

	f, err := ioutil.ReadFile("./chain.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))

	err = json.Unmarshal(f, &ThisChain)
	if err != nil {
		panic(err)
	}

	/*

		for _, peer := range peers {
			if fetch(peer) == nil {
				break
			}
		}

		if len(peers) == 0 {
			for _, peer := range PermanentPeers {
				fmt.Println(peer)
				if fetch(peer) == nil {
					break
				}
			}
		}*/
}

func AddTransactionToMempool(transaction Transaction) {
	if transaction.IsFieldEmpty() || !transaction.VerifySignature() {
		return
	}

	Mempool = append(Mempool, transaction)
}

func AddBlock(block Block) error {
	if !block.Verify() {
		return errors.New("invalid block")
	}

	if !bytes.Equal(block.Header.PreviousHash[:], ThisChain[len(ThisChain)-1].Header.Hash[:]) {
		return errors.New("invalid previous hash")
	}

	ThisChain = append(ThisChain, block)

	err := WriteChain(ThisChain)
	if err != nil {
		return err
	}

	Mempool = []Transaction{}

	return nil
}

func WriteChain(chain Chain) error {
	j, err := json.MarshalIndent(ThisChain, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./chain.json", j, 0644)
}

func (c Chain) CalculateTarget() []byte {
	z := new(big.Int)
	z.SetBytes(Max)

	target := new(big.Int).Div(z, big.NewInt(1_000_000))
	return target.Bytes()
}

func (c Chain) CalculateReward() uint64 {
	return FractionsPerCoin * 10
}

func ValidNonce(nonce uint64) bool {
	for _, block := range ThisChain {
		for _, tx := range block.Transactions {
			if tx.Data.Nonce == nonce {
				return false
			}
		}
	}

	return true
}

func Balance(publicKey PublicKey) uint64 {
	var balance uint64

	for _, block := range ThisChain {
		for _, tx := range block.Transactions {
			if tx.Data.Receiver == publicKey {
				balance += tx.Data.Amount
			} else if tx.Data.Sender == publicKey {
				balance -= tx.Data.Amount
			}
		}
	}

	return balance
}
