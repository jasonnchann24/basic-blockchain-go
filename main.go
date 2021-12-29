package main

import (
	"fmt"
	"strconv"

	"github.com/jasonnchann24/learnBlockChain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for i, block := range chain.Blocks {
		fmt.Println()
		fmt.Printf("Block %d\n", i)
		fmt.Println("........")
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("============================================================")
	}
}
