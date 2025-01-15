package transaction

import "bytes"

type TxInput struct {
	TxID        []byte
	OutIdx      int
	FromAddress []byte
}

type TxOutput struct {
	Value     int
	ToAddress []byte
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.FromAddress, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.ToAddress, address)
}
