package core

func (c Chain) GetBalance(publicKey PublicKey, regardMempool bool) uint64 {
	var balance uint64
	if regardMempool {
		for _, tx := range Mempool {
			if tx.Data.Sender.Equals(publicKey) {
				balance = balance - tx.Data.Amount
			}

			if tx.Data.Receiver.Equals(publicKey) {
				balance = balance + tx.Data.Amount
			}
		}
	}

	for _, block := range c {
		for _, tx := range block.Transactions {
			if tx.Data.Sender.Equals(publicKey) {
				balance = balance - tx.Data.Amount
			}

			if tx.Data.Receiver.Equals(publicKey) {
				balance = balance + tx.Data.Amount
			}
		}
	}

	balance = balance + (FractionsPerCoin / 10)

	return balance
}
