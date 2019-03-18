package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18545")
	if err != nil {
		log.Fatal(err)
	}

	// get specific block
	blockNumber := big.NewInt(5181405)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// get transactions information
	for index, tx := range block.Transactions() {
		// https://ropsten.etherscan.io/tx/0x412187a25b1464fdcb6ad3e08844701a67f2f3b3c5751be1b21aea901ed55cad
		fmt.Printf("\nTransaction #%d\n",index) // Transaction #24
		fmt.Println(tx.Hash().Hex())        // 0x412187a25b1464fdcb6ad3e08844701a67f2f3b3c5751be1b21aea901ed55cad
		fmt.Println(tx.Value().String())    // 1000000000000000000
		fmt.Println(tx.Gas())               // 21000
		fmt.Println(tx.GasPrice().Uint64()) // 1000000000
		fmt.Println(tx.Nonce())             // 22342721
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x2E2F47e77A2ebFceB519155470b3F06a2BFE06f8

		// get sender address
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		// TODO chainID?
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
			fmt.Println(msg.From().Hex()) // 0x81b7E08F65Bdf5648606c89998A9CC8164397647
		}

		// each transaction has a receipt which contains the result of the execution of the transaction
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receipt from last transaction: %d\n\n", receipt.Status) // 1
	}

	// instead of iterate over transaction in block
	// we can call TransactionInBlock
	// require block hash and index of the transaction within the block
	blockHash := common.HexToHash("0x288c9747d5d3007b5c3cb53a90ead5217c34040ec7875f315c9d875646a212ed")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nGet tx from blockhash and tx index")
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}

	// query for a single transaction directly
	fmt.Println("\nQuery for a single transaction directly")
	txHash := common.HexToHash("0x412187a25b1464fdcb6ad3e08844701a67f2f3b3c5751be1b21aea901ed55cad")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex()) // 0x412187a25b1464fdcb6ad3e08844701a67f2f3b3c5751be1b21aea901ed55cad
	fmt.Println(isPending)       // false
}