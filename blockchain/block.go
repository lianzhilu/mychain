//block.go

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"github.com/lianzhilu/mychain/utils"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{utils.Int64ToByte(b.Timestamp), b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func CreateBlock(prevHash, data []byte) *Block {
	block := Block{
		Timestamp: time.Now().Unix(),
		Hash:      []byte{},
		PrevHash:  prevHash,
		Data:      data,
	}
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	genesisWords := "HelloWorld"
	return CreateBlock([]byte{}, []byte(genesisWords))
}
