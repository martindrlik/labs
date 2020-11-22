package account_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martindrlik/labs/account"
)

const amanda = "amanda"

func ExampleIsAvalilable() {
	as := new(account.Accounts)
	fmt.Println(as.IsAvailable(amanda))
	as.Add(map[string]account.Account{amanda: {}})
	fmt.Println(as.IsAvailable(amanda))
	// Output:
	// true
	// false
}

func ExampleDeactivate() {
	as := new(account.Accounts)
	as.Add(map[string]account.Account{amanda: {}})
	if as.Deactivate(amanda) {
		fmt.Println("Amanda's account is deactivated")
		fmt.Printf("Amanda's name is not available: %v\n", !as.IsAvailable(amanda))
	}
	// Output:
	// Amanda's account is deactivated
	// Amanda's name is not available: true
}

const amandakey = "amanda + amanda's password hash"

func ExampleIsValid() {
	ks := new(account.Keys)
	fmt.Println(ks.IsValid(amandakey))
	ks.Add(amandakey)
	fmt.Println(ks.IsValid(amandakey))
	// Output:
	// false
	// true
}

func ExampleInvalidate() {
	ks := new(account.Keys)
	ks.Add(amandakey)
	if ks.Invalidate(amandakey) {
		fmt.Printf("Amanda's key is no longer valid: %v\n", !ks.IsValid(amandakey))
	}
	// Output:
	// Amanda's key is no longer valid: true
}

func ExampleKeys() {
	ks := new(account.Keys)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(ks.Keys())
	ks.Add(amandakey)
	enc.Encode(ks.Keys())
	// Output:
	// []
	// ["amanda + amanda's password hash"]
}

func ExampleAccounts() {
	as := new(account.Accounts)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(as.Accounts())
	as.Add(map[string]account.Account{amanda: {}})
	enc.Encode(as.Accounts())
	// Output:
	// []
	// [{"Name":"amanda","Inactive":false}]
}
