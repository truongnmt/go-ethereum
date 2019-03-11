package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18545")
	if err != nil {
		log.Fatalf("Failed to call ethclient.Dial()\n%v", err)
	}

	// setting nil at blockNumber will return latest balance
	account := common.HexToAddress("0x2E2F47e77A2ebFceB519155470b3F06a2BFE06f8")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("Failed to call client.BalanceAt()\n%v", err)
	}
	fmt.Printf("Latest balance: %v\n", balance) // 2000000000000000000

	// passing blockNumber will return account balance at the time of that block
	blockNumber := big.NewInt(5181337)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatalf("Failed to call client.BalanceAt()\n%v", err)
	}
	fmt.Printf("Balance at block 5181337: %v\n", balanceAt) // 1000000000000000000


	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}