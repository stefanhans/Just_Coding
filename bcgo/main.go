package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Soumik says hi")
	bc.AddBlock("and Stefan also")

	fmt.Printf("\nPrint the blockchain:\n----------------------\n\n")

	for _, block := range bc.blocks {
		fmt.Printf("Previous Hash: \t%x\n", block.PrevBlockHash)
		fmt.Printf("Data: \t\t%s\n", block.Data)
		fmt.Printf("Current Hash: \t%x\n", block.Hash)
		fmt.Println()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: \t\t%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
