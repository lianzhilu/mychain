package blockchain

import (
	"encoding/hex"
	"fmt"
	"github.com/lianzhilu/mychain/transaction"
)

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(txs []*transaction.Transaction) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, txs)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}

func (bc *BlockChain) FindUnspentTransactions(address []byte) []transaction.Transaction {
	var unSpentTxs []transaction.Transaction      // 用于存储未花费的交易
	spentTxs := make(map[string]map[int]struct{}) // 用于跟踪已消费的交易输出，结构为 {txID: {outputIndex: struct{}}}

	// 从最新区块向前遍历
	for idx := len(bc.Blocks) - 1; idx >= 0; idx-- {
		block := bc.Blocks[idx]

		// 遍历每个区块中的交易
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

			// 遍历每个交易的输出
			for outIdx, out := range tx.Outputs {
				// 如果当前输出已经被消费，则跳过
				if isOutputSpent(txID, outIdx, spentTxs) {
					continue
				}

				// 如果输出属于目标地址，则记录为未花费
				if out.ToAddressRight(address) {
					unSpentTxs = append(unSpentTxs, *tx)
				}
			}

			// 如果交易是非Base交易，记录其输入，标记已消费的输出
			if !tx.IsBase() {
				for _, in := range tx.Inputs {
					if in.FromAddressRight(address) {
						inTxID := hex.EncodeToString(in.TxID)
						if spentTxs[inTxID] == nil {
							spentTxs[inTxID] = make(map[int]struct{})
						}
						spentTxs[inTxID][in.OutIdx] = struct{}{}
					}
				}
			}
		}
	}

	return unSpentTxs
}

// isOutputSpent 检查某个交易输出是否已被消费
func isOutputSpent(txID string, outIdx int, spentTxs map[string]map[int]struct{}) bool {
	if spentOutputs, exists := spentTxs[txID]; exists {
		if _, spent := spentOutputs[outIdx]; spent {
			return true
		}
	}
	return false
}

func (bc *BlockChain) FindUTXOs(address []byte) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) {
				accumulated += out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) FindSpendableOutputs(address []byte, amount int) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				if accumulated >= amount {
					break Work
				}
				continue Work
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) CreateTransaction(from, to []byte, amount int) (*transaction.Transaction, bool) {
	var inputs []transaction.TxInput
	var outputs []transaction.TxOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		fmt.Println("Not enough coins!")
		return &transaction.Transaction{}, false
	}
	for txid, outidx := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			panic(err)
		}
		input := transaction.TxInput{txID, outidx, from}
		inputs = append(inputs, input)
	}

	outputs = append(outputs, transaction.TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, transaction.TxOutput{acc - amount, from})
	}
	tx := transaction.Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx, true
}

func (bc *BlockChain) Mine(txs []*transaction.Transaction) {
	bc.AddBlock(txs)
}
