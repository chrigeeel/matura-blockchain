package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type Block struct {
	MagicNumber        uint64        `json:"magicNumber"`
	TransactionCounter uint64        `json:"transactionCounter"`
	Transactions       []Transaction `json:"transactions"`
	Header             BlockHeader   `json:"header"`
}

type BlockHeader struct {
	Version      string    `json:"version"`
	PreviousHash Hash      `json:"previousHash"`
	Miner        PublicKey `json:"publicKey"`
	Solution     []byte    `json:"solution"`
	Hash         Hash      `json:"hash"`
	Time         int64     `json:"time"`
}

const blockchainMagicNumber = 0xD9B4BEF9

func BuildBlock(pub PublicKey, solution []byte) (Block, error) {
	thisMempool := Mempool

	var validTransactions []Transaction

	for _, tx := range thisMempool {
		fmt.Println(tx)
		if tx.VerifyTotal(true) {
			fmt.Println("passed")
			validTransactions = append(validTransactions, tx)
		}
	}

	validTransactions = append(validTransactions, Transaction{
		IsCoinbase: true,
		Data: TransactionData{
			Receiver: pub,
			Amount:   ThisChain.CalculateReward(),
		},
	})

	block := Block{
		MagicNumber:        blockchainMagicNumber,
		TransactionCounter: uint64(len(validTransactions)),
		Transactions:       validTransactions,
		Header: BlockHeader{
			Version:      Version,
			PreviousHash: ThisChain[len(ThisChain)-1].Header.Hash,
			Miner:        pub,
			Solution:     solution,
			Time:         time.Now().Unix(),
		},
	}

	block.Header.Hash = block.Hash()

	return block, nil
}

func (b Block) Hash() Hash {
	b.Header.Hash = Hash{}
	j, _ := json.Marshal(b.Header)
	return HashBytes(j)
}

func (b Block) Verify(regardMempool bool) bool {
	return b.VerifyBlockMiner() && b.VerifyBlockSolution() && b.VerifyBlockTransactions(regardMempool)
}

func (b Block) VerifyBlockSolution() bool {
	target := ThisChain.CalculateTarget()

	for _, block := range ThisChain {
		if bytes.Equal(block.Header.Solution, b.Header.Solution) {
			return false
		}
	}

	return VerifySolution(target, b.Header.Solution)
}

func (b Block) VerifyBlockMiner() bool {
	return strings.HasPrefix(string(b.Header.Solution), b.Header.Miner.String())
}

func (b Block) VerifyBlockTransactions(regardMempool bool) bool {
	var rewardPaid bool
	for _, tx := range b.Transactions {
		if tx.IsCoinbase {
			if tx.Data.Receiver != b.Header.Miner || rewardPaid || tx.Data.Amount != ThisChain.CalculateReward() {
				return false
			}
			rewardPaid = true
			continue
		}

		if !tx.VerifyTotal(regardMempool) {
			return false
		}
	}
	return true
}

func VerifySolution(target []byte, solution []byte) bool {
	targetBigInt := new(big.Int)
	targetBigInt.SetBytes(target)

	solutionBigInt := new(big.Int)
	hash := HashBytes(solution)
	solutionBigInt.SetBytes(hash[:])

	return targetBigInt.Cmp(solutionBigInt) == 1
}
