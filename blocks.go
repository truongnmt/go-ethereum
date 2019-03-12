package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18545")
	if err != nil {
		log.Fatal(err)
	}

	// latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5186255

	// Get full block data
	blockNumber := big.NewInt(5181337)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5181337
	fmt.Println(block.Time().Uint64())       // 1552291080
	fmt.Println(block.Difficulty().Uint64()) // 2250650754
	fmt.Println(block.Hash().Hex())          // 0xcb3ce38b35914c6f6e3fc140c3f8ac52c072d6a1b4fda24a99306f7258913134
	fmt.Println(len(block.Transactions()))   // 73 list of transaction
	//for _, tx := range block.Transactions() {
	//	fmt.Println(tx)
	//}

	// call TransactionCount to return only the count of transactions in a block
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 73
}