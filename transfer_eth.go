package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mf-financial/wallet-util/logger"
)

func main() {
	client, err := ethclient.Dial("http://localhost:18545")
	if err != nil {
		logger.ErrorWithStack(err)
	}

	// load private key
	privateKey, err := crypto.HexToECDSA("636D5104B58403E85C038070932916C9CB737C5E61D2BBD94500E46186430BBE")
	if err != nil {
		logger.ErrorWithStack(err)
	}

	// get the account nonce
	// nonce is only used once
	// TODO don't forget the Public() part, or we will send private key instead
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	// https://etherconverter.online
	value := big.NewInt(500000000000000000) // in wei (0.5 eth)
	gasLimit := uint64(21000)                   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.ErrorWithStack(err)
	}
	fmt.Printf("Gas limit: %v\n", gasLimit)
	fmt.Printf("Gas price: %v\n", gasPrice)

	toAddress := common.HexToAddress("0x2E2F47e77A2ebFceB519155470b3F06a2BFE06f8")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.ErrorWithStack(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}