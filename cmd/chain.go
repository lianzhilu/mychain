package chain

import (
	"fmt"
	"github.com/lianzhilu/mychain/blockchain"
	"github.com/spf13/cobra"
	"time"
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
	bc := blockchain.CreateBlockChain()
	time.Sleep(time.Second)
	bc.AddBlock("first")
	time.Sleep(time.Second)
	bc.AddBlock("second")
	time.Sleep(time.Second)
	bc.AddBlock("third")
	time.Sleep(time.Second)

	for num, block := range bc.Blocks {
		fmt.Printf("number:%d Timestamp: %d\n", num, block.Timestamp)
		fmt.Printf("number:%d hash: %x\n", num, block.Hash)
		fmt.Printf("number:%d Previous hash: %x\n", num, block.PrevHash)
		fmt.Printf("number:%d data: %s\n", num, block.Data)
		fmt.Println("POW validation:", block.ValidatePoW())
	}
}
