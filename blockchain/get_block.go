package blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/onflow/flow-go-sdk"
	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func GetBlock() {

	ctx := context.Background()
	flowClient := flowInstance.NewFlowClient()

	// Get the latest sealed block
	isSealed := true
	latestBlock, err := flowClient.GetLatestBlock(ctx, isSealed)
	if err != nil {
		log.Fatal(err)
	}
	printBlock(latestBlock, err)

	// get the block by id
	blockId := latestBlock.ID.String()
	blockById, err := flowClient.GetBlockByID(ctx, flow.HexToID(blockId))
	if err != nil {
		log.Fatal(err)
	}
	printBlock(blockById, err)

	// get block by height
	blockByHeight, err := flowClient.GetBlockByHeight(ctx, 0)
	if err != nil {
		log.Fatal(err)
	}
	printBlock(blockByHeight, err)

}

func printBlock(block *flow.Block, err error) {
	// flowClient := flowInstance.NewFlowClient()

	// flowClient.Handle(err)

	fmt.Printf("\nID: %s\n", block.ID)
	fmt.Printf("height: %d\n", block.Height)
	fmt.Printf("timestamp: %s\n\n", block.Timestamp)
}
