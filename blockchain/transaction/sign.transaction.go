package transaction

import (
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

var (
	myAddress    flow.Address
	myAccountKey flow.AccountKey
	myPrivateKey crypto.PrivateKey
)

func SignTransaction() *flow.Transaction {
	tx := flow.NewTransaction().
		SetScript([]byte("transaction { execute { log(\"Hello, World!\") } }")).
		SetGasLimit(100).
		SetProposalKey(myAddress, myAccountKey.Index, myAccountKey.SequenceNumber).
		SetPayer(myAddress)

	return tx
}
