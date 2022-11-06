package core

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"
)

func Mine(k PrivateKey) {
	var hashPerSecond int
	target := ThisChain.CalculateTarget()
	publicKey := k.PublicKey()
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			log.Printf("Mining... | %d H/s\n", hashPerSecond)
			hashPerSecond = 0
		}
	}()
	for {
		nonce, err := RandomNonce()
		if err != nil {
			continue
		}
		solution := fmt.Sprintf("%s-%s", publicKey.String(), nonce)
		hashPerSecond++
		if !VerifySolution(target, []byte(solution)) {
			continue
		}

		log.Printf("💀 | Mined new block! | Hash: %s \n", HashString(solution).String())
		block, err := BuildBlock(k.PublicKey(), []byte(solution))
		if err != nil {
			panic(err)
		}
		err = BroadcastBlock(block)
		if err != nil {
			panic(err)
		}
	}
}

func RandomNonce() (string, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	return n.String(), nil
}
