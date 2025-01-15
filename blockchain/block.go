//block.go

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"github.com/lianzhilu/mychain/transaction"
	"github.com/lianzhilu/mychain/utils"
	"math/big"
	"time"
)

type Block struct {
	Timestamp    int64                      // 区块被创建的是时间
	Hash         []byte                     // 区块大哈希值
	PrevHash     []byte                     // 上一个区块的哈希值
	Nonce        int64                      // 随机数, 用于PoW
	Target       []byte                     // 目标值, 用于验证PoW
	Transactions []*transaction.Transaction // 区块中包含的交易
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{
		utils.Int64ToByte(b.Timestamp),
		b.PrevHash,
		b.Target,
		utils.Int64ToByte(b.Nonce),
		b.BackTXSummary(),
	}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func CreateBlock(prevHash []byte, txs []*transaction.Transaction) *Block {
	block := Block{
		Timestamp:    time.Now().Unix(),
		Hash:         []byte{},
		PrevHash:     prevHash,
		Nonce:        0,
		Target:       []byte{},
		Transactions: txs,
	}
	block.InitPoW()
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	tx := transaction.BaseTx([]byte("Base"))
	return CreateBlock([]byte{}, []*transaction.Transaction{tx})
}

func (b *Block) ValidatePoW() bool {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	intTarget.SetBytes(b.Target)
	data := b.GetDataBaseNonce(b.Nonce)
	hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	if intHash.Cmp(&intTarget) == -1 {
		return true
	}
	return false
}

func (b *Block) BackTXSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	summary := bytes.Join(txIDs, []byte{})
	return summary
}
