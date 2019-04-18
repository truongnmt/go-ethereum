package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// connect to an ethereum node or infura
	client, err := ethclient.Dial("http://localhost:18545")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := NewInbox(common.HexToAddress("0x138d331A9837c266d2764CfD5f217d94F9cf9daE"), client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	fmt.Println("Contract getMessage():")
	fmt.Println(contract.Message(nil))
}
