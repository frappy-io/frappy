package blockchain

import (
	"context"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func GetCollection(collectionId flow.Identifier) {

	ctx := context.Background()
	flowClient := flowInstance.NewFlowClient()

	// get collection by ID
	collection, err := flowClient.GetCollection(ctx, collectionId)
	printCollection(collection, err)
}

func printCollection(collection *flow.Collection, err error) {

	fmt.Printf("\nID: %s", collection.ID().String())
	fmt.Printf("\nTransactions: %s", collection.TransactionIDs)
}
