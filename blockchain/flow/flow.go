package flow

import "github.com/onflow/flow-go-sdk/client"

// common client Interface
var flowClient client.Client

func NewFlowClient() client.Client {
	return flowClient
}
