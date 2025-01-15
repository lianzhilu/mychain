package main

import (
	"fmt"
	"github.com/lianzhilu/mychain/blockchain"
	"time"
)

func main() {
	bc := blockchain.CreateBlockChain()
	time.Sleep(time.Second)
	bc.AddBlock("This is first Block after Genesis")
	time.Sleep(time.Second)
	bc.AddBlock("This is second!")
	time.Sleep(time.Second)
	bc.AddBlock("Awesome!")
	time.Sleep(time.Second)

	for num, block := range bc.Blocks {
		fmt.Printf("number:%d Timestamp: %d\n", num, block.Timestamp)
		fmt.Printf("number:%d hash: %x\n", num, block.Hash)
		fmt.Printf("number:%d Previous hash: %x\n", num, block.PrevHash)
		fmt.Printf("number:%d data: %s\n", num, block.Data)
	}
}
