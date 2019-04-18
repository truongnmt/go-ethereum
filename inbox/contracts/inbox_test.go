package contracts

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestDeployInbox(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10)

	//Deploy contract
	address, _, _, err := DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)
	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the Inbox contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

}

//Test initial message gets set up correctly
func TestGetMessage(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10)

	//Deploy contract
	_, _, contract, _ := DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	// commit all pending transactions
	blockchain.Commit()

	if got, _ := contract.Message(nil); got != "Hello World" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}

// Test message gets updated correctly
func TestSetMessage(t *testing.T) {

	//Setup simulated blockchain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10)

	//Deploy contract

	_, _, contract, _ := DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	// commit all pending transactions
	blockchain.Commit()
	contract.SetMessage(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, "Hello from Mars")

	blockchain.Commit()

	if got, _ := contract.Message(nil); got != "Hello from Mars" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}