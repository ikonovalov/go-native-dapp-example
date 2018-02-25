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
	"context"
	"time"
	"github.com/ethereum/go-ethereum/common"
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
	fmt.Printf("Contract   address: %s\n", addr.String())

	mined := make(chan common.Address)
	ctx := context.Background()

	go func() {
		fmt.Println("Wait deployed...")
		address, _ := bind.WaitDeployed(ctx, sim, tx)
		mined <- address
		close(mined)
	}()

	sim.Commit()

	select {
	case a := <-mined:
		fmt.Printf("Mined! %s\n", a.String())
	case <-time.After(20 * time.Second):
		fmt.Errorf("timeout")
	}

	opts := auth
	tx, err = contract.Greet(opts, "h2")
	fmt.Printf("tx=%s\n", tx.Hash().String())
	sim.Commit()
	name, err := contract.Name(nil)
	cnt, err := contract.Count(nil)
	fmt.Println("Name=", name)
	fmt.Println("Count=", cnt)

}
