package blockchain

import (
	"bytes"
	"crypto/sha256"
	"github.com/lianzhilu/mychain/utils"
	"math"
	"math/big"
)

func (b *Block) GetTarget() {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-utils.Difficulty))
	b.Target = target.Bytes()
}

func (b *Block) GetDataBaseNonce(nonce int64) []byte {
	data := bytes.Join([][]byte{
		utils.Int64ToByte(b.Timestamp),
		b.PrevHash,
		utils.Int64ToByte(nonce),
		b.Target,
		b.Data,
	},
		[]byte{},
	)
	return data
}

func (b *Block) FindNonce() {
	var intHash big.Int
	var intTarget big.Int

	intTarget.SetBytes(b.Target)

	var hash [32]byte
	var nonce int64
	nonce = 0

	for nonce < math.MaxInt64 {
		data := b.GetDataBaseNonce(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(&intTarget) == -1 {
			break
		} else {
			nonce++
		}
	}

	b.Nonce = nonce
}

func (b *Block) InitPoW() {
	b.GetTarget()
	b.FindNonce()
}
