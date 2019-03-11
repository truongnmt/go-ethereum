package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// generate a random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// private key convert to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // d593ae0363f9e66fe42a91fe0260067439e0760a3f4f554696cd9350c41ae7a6

	// private key to public key
	publicKey := privateKey.Public()

	// public key to hex
	// slice the 0x and the first 2 character 04 which is always the EC prefix and not required
	// interface typecast
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // b00316b00f432d90c1a3b5ee0f0fc4416c5aab07677edcc0ca65c53eb100f92fbde998d466bf7676cb4eaa2bf78f1a1df086718d54170e47b6d59315002d61df

	// generate public address from public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x2Fc75Ce9cE3529758b5787ADE183566F765De0b9

	// to hex
	// TODO (?)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x2fc75ce9ce3529758b5787ade183566f765de0b9
}