package account_test

import (
	"fmt"

	"github.com/martindrlik/labs/account"
)

func ExampleAccounts() {
	const amanda = "amanda"
	as := new(account.Accounts)
	fmt.Println(as.IsAvailable(amanda))
	as.Add(map[string]account.Account{amanda: {}})
	fmt.Println(as.IsAvailable(amanda))
	if as.Deactivate(amanda) {
		fmt.Println("Amanda's account deactivated")
	}
	fmt.Println(as.IsAvailable(amanda))
	// Output:
	// true
	// false
	// Amanda's account deactivated
	// false
}

func ExampleKeys() {
	const amandakey = "amanda + amanda's password hash"
	ks := new(account.Keys)
	fmt.Println(ks.IsValid(amandakey))
	ks.Add(amandakey)
	fmt.Println(ks.IsValid(amandakey))
	ks.Invalidate(amandakey)
	fmt.Println(ks.IsValid(amandakey))
	// Output:
	// false
	// true
	// false
}
