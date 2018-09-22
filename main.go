package main

import (
	"fmt"
    "strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Minsoo에게 1 GoGoCoin을 전송합니다.")
	bc.AddBlock("Jisoo가 2 GoGoCoin으로 피자를 구매했습니다.")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
