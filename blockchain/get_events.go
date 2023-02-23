package blockchain

import (
	// "context"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	// flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func GetEvents() {
	// ctx := context.Background()
	// flowClient := flowInstance.NewFlowClient()

	// Query for account creation events by type
	// result, err := flowClient.GetEventsForHeightRange(ctx, "flow.AccountCreated", 0, 30)
	// printEvents(result, err)

	// Query for our custom event by type
	// customType := fmt.Sprintf("AC.%s.EventDemo.EventDemo.Add", deployedContract.Address.Hex())
	// result, err = flowClient.GetEventsForHeightRange(ctx, customType, 0, 10)
	// printEvents(result, err)

	// Get events directly from transaction result
	// txResult, err := flowClient.GetTransactionResult(ctx, runScriptTx.ID())
	// examples.Handle(err)
	// printEvent(txResult.Events)
}

func printEvents(result []client.BlockEvents, err error) {

	for _, block := range result {
		printEvent(block.Events)
	}
}

func printEvent(events []flow.Event) {
	for _, event := range events {
		fmt.Printf("\n\nType: %s", event.Type)
		fmt.Printf("\nValues: %v", event.Value)
		fmt.Printf("\nTransaction ID: %s", event.TransactionID)
	}
}
