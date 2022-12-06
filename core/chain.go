package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Chain []Block

var (
	ThisChain Chain
)

func AddBlock(block Block, regardMempool bool) error {
	if !block.Verify(regardMempool) {
		return errors.New("invalid block")
	}

	if !bytes.Equal(block.Header.PreviousHash[:], ThisChain[len(ThisChain)-1].Header.Hash[:]) {
		return errors.New("invalid previous hash")
	}

	fmt.Println("adding block")

	ThisChain = append(ThisChain, block)

	err := WriteChain()
	if err != nil {
		return err
	}

	ClearMempool()

	return nil
}

func WriteChain() error {
	j, err := json.MarshalIndent(ThisChain, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./chain.json", j, 0644)
}

func (c Chain) CalculateReward() uint64 {
	return FractionsPerCoin * 10
}
