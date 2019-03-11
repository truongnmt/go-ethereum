package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18545")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("we have a connection")
	_ = client
}
