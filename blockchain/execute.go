package blockchain

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/examples"
	flowInstance "github.com/siddhantprateek/frappy/blockchain/flow"
)

func ExecuteScript() {
	ctx := context.Background()
	flowClient := flowInstance.NewFlowClient()

	script := []byte(`
		pub fun main(a: Int) Int {
			return a + 10
		}
	`)
	args := []cadence.Value{cadence.NewInt(5)}
	value, err := flowClient.ExecuteScriptAtLatestBlock(ctx, script, args)

	examples.Handle(err)
	fmt.Printf("\nValue: %s", value.String())

	complexScript := []byte(`
		pub struct User {
			pub var balance: UFix64
			pub var address: Address
			pub var name: String

			init(name: String, address: Address, balance: UFix64){
				self.name = name
				self.address = address
				self.balance = balance
			}
		}

		pub fun main(name: String): User {
			return User(
				name: name,
				address: 0x1,
				balance: 10.0
			)
		}
	`)
	// args = []cadence.Value{cadence.NewString("Dete")}
	args = []cadence.Value{}
	value, err = flowClient.ExecuteScriptAtLatestBlock(ctx, complexScript, args)
	printComplexScript(value, err)
}

type User struct {
	balance uint64
	address flow.Address
	name    string
}

func printComplexScript(value cadence.Value, err error) {
	examples.Handle(err)
	fmt.Printf("\nString value: %s", value.String())

	s := value.(cadence.Struct)
	u := User{
		balance: s.Fields[0].ToGoValue().(uint64),
		address: s.Fields[1].ToGoValue().([flow.AddressLength]byte),
		name:    s.Fields[2].ToGoValue().(string),
	}

	fmt.Printf("\nName: %s", u.name)
	fmt.Printf("\nAddress: %s", u.address.String())
	fmt.Printf("\nBalance: %d", u.balance)
}
