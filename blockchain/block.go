//block.go

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"github.com/lianzhilu/mychain/utils"
	"math/big"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     int64
	Target    []byte
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
		Nonce:     0,
		Target:    []byte{},
	}
	block.InitPoW()
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	genesisWords := "HelloWorld"
	return CreateBlock([]byte{}, []byte(genesisWords))
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
