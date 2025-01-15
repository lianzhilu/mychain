package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/lianzhilu/mychain/utils"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func (tx *Transaction) CalculateTXHash() []byte {
	var buffer bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		panic(err)
	}

	hash = sha256.Sum256(buffer.Bytes())
	return hash[:]
}

func (tx *Transaction) SetID() {
	tx.ID = tx.CalculateTXHash()
}

func BaseTx(toAddress []byte) *Transaction {
	txIn := TxInput{[]byte{}, -1, []byte{}}
	txOut := TxOutput{utils.InitCoin, toAddress}
	tx := Transaction{[]byte("This is the Base Transaction!"),
		[]TxInput{txIn}, []TxOutput{txOut}}
	return &tx
}

func (tx *Transaction) IsBase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutIdx == -1
}
