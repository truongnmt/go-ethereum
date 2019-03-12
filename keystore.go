package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	// keystore contains an encrypted wallet private key
	// keystore contain one wallet key pair per file
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x773E649E140a71d84bb1b24eCd36541587474C99
}

func importKs() {
	file := "./tmp/UTC--2019-03-12T01-24-32.718972000Z--773e649e140a71d84bb1b24ecd36541587474c99"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// import keystore from file
	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex()) // 0x773E649E140a71d84bb1b24eCd36541587474C99

	// TODO every time import will generate a new keystore file, so we delete the old one (weird)
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}