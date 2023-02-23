package blockchain

import (
	"context"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

// Retrieve any account from Flow network's latest block

func GetAccount() {

	ctx := context.Background()
	flowClient := flowInstance.NewFlowClient()

	// get account from the latest block
	address := flow.HexToAddress("f8d6e0586b0a20c7")
	account, err := flowClient.GetAccount(ctx, address)
	printAccount(account, err)

	// get account from the block by height 0
	account, err = flowClient.GetAccountAtBlockHeight(ctx, address, 0)
	printAccount(account, err)

}

func printAccount(account *flow.Account, err error) {

	fmt.Printf("\nAddress: %s", account.Address.String())
	fmt.Printf("\nBalance: %d", account.Balance)
	fmt.Printf("\nContracts: %d", len(account.Contracts))
	fmt.Printf("\nKeys: %d\n", len(account.Keys))
}
