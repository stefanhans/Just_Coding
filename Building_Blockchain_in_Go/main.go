package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Soumik says hi")
	bc.AddBlock("And Stefan as well")

	for _, block := range bc.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
	}
}
