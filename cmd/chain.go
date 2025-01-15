package chain

import (
	"github.com/lianzhilu/mychain/blockchain"
	"github.com/lianzhilu/mychain/transaction"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			RunChain()
		},
	}
	return cmd
}

func RunChain() {
	txPool := make([]*transaction.Transaction, 0)
	var tempTx *transaction.Transaction
	var ok bool
	chain := blockchain.CreateBlockChain()

	tempTx, ok = chain.CreateTransaction([]byte("first"), []byte("second"), 100)
	if ok {
		txPool = append(txPool, tempTx)
	}
	chain.Mine(txPool)
}
