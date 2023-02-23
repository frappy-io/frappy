package blockchain

import (
	"context"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func GetTransaction(txId flow.Identifier) {
	ctx := context.Background()
	flowClient := flowInstance.NewFlowClient()

	tx, err := flowClient.GetTransaction(ctx, txId)
	printTransaction(tx, err)

	txr, err := flowClient.GetTransactionResult(ctx, txId)
	printTransactionResult(txr, err)
}

func printTransaction(tx *flow.Transaction, err error) {

	fmt.Printf("\nID: %s", tx.ID().String())
	fmt.Printf("\nPayer: %s", tx.Payer.String())
	fmt.Printf("\nProposer: %s", tx.ProposalKey.Address.String())
	fmt.Printf("\nAuthorizers: %s", tx.Authorizers)
}

func printTransactionResult(txr *flow.TransactionResult, err error) {

	fmt.Printf("\nStatus: %s", txr.Status.String())
	fmt.Printf("\nError: %v", txr.Error)
}
