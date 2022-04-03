package contracts

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// Test inbox contract gets deployed correctly
func TestDeployInbox(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()

	transactOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	contractBackend := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			transactOpts.From: {Balance: big.NewInt(10000000000000000)},
		}, 10000000,
	)
	defer contractBackend.Close()

	//Deploy contract
	address, _, contract, err := DeployInbox(
		transactOpts,
		contractBackend,
		"Hello World",
	)
	// commit all pending transactions
	contractBackend.Commit()

	// Test Deploy
	if err != nil {
		t.Fatalf("Failed to deploy the Inbox contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

	// Test GettMessage
	if got, _ := contract.Message(nil); got != "Hello World" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}

	// Test SetMessage
	contract.SetMessage(&bind.TransactOpts{
		From:   transactOpts.From,
		Signer: transactOpts.Signer,
		Value:  nil,
	}, "Hello from Mars")

	contractBackend.Commit()

	if got, _ := contract.Message(nil); got != "Hello from Mars" {
		t.Errorf("Expected message to be: Hello World. Go: %s", got)
	}
}
