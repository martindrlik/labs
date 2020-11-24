package account_test

import (
	"bytes"
	"fmt"

	"github.com/martindrlik/labs/account"
)

func ExampleAccountStore() {
	const a = "a"
	salt := []byte("salt")
	s := new(account.Store)
	if s.Add(map[string]account.Account{a: {Active: true, Salt: salt}}) {
		fmt.Println("a is added")
	}
	if s.Account(a).Active {
		fmt.Println("a is active")
	}
	if bytes.Equal(s.Account(a).Salt, salt) {
		fmt.Println("a salt is salt")
	}
	if s.Deactivate(a) {
		fmt.Println("a is deactivated")
	}
	// Output:
	// a is added
	// a is active
	// a salt is salt
	// a is deactivated
}
