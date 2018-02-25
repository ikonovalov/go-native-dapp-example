package main

import (
	"github.com/ikonovalov/go-native-dapp-example/contracts/gen"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc)

	addr, tx, contract, err := greeter.DeployGreeter(auth, sim)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	fmt.Printf("Deploy transaction: %s\n", tx.Hash().String())
	fmt.Printf("Contract   address: %s\n", addr.Hash().String())
	fmt.Println("Mining...")
	sim.Commit()

	opts := auth
	tx, err = contract.Greet(opts, "h2")
	fmt.Printf("tx=%s\n", tx.Hash().String())
	sim.Commit()
	name, err := contract.Name(nil)
	cnt, err := contract.Count(nil)
	fmt.Println("Name=", name)
	fmt.Println("Count=", cnt)

}
