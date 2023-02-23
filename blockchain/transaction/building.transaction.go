package transaction

import (
	"context"
	"os"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func BuildTransaction() {

	greeting, err := os.ReadFile("Greeting2.cdc")
	if err != nil {
		panic("failded to load Cadence script")
	}

	proposerAddress := flow.HexToAddress("9a0766d93b6608b7")
	proposerKeyIndex := 3

	payerAddress := flow.HexToAddress("631e88ae7f1d7c20")
	authorizerAddress := flow.HexToAddress("7aad92e5a0715d21")

	// var accessAPIHost string

	// Establish a connection with an access node
	flowClient := flowInstance.NewFlowClient()

	// Get the latest sealed block to use as a reference block
	latestBlock, err := flowClient.GetLatestBlockHeader(context.Background(), true)
	if err != nil {
		panic("failded to fetch proposer account")
	}

	// Get the latest sequence number for this key
	proposerAccount, err := flowClient.GetAccountAtLatestBlock(context.Background(), proposerAddress)
	if err != nil {
		panic("failed to fetch proposer account")
	}

	// Get the latest sequence number for this key
	sequenceNumber := proposerAccount.Keys[proposerKeyIndex].SequenceNumber

	tx := flow.NewTransaction().
		SetScript(greeting).
		SetGasLimit(100).
		SetReferenceBlockID(latestBlock.ID).
		SetProposalKey(proposerAddress, proposerKeyIndex, sequenceNumber).
		SetPayer(payerAddress).
		AddAuthorizer(authorizerAddress)

	// Add arguments last

	hello, _ := cadence.NewString("Hello")
	err = tx.AddArgument(hello)
	if err != nil {
		panic("invalid argument")
	}
}
