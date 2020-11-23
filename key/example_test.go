package key_test

import (
	"bytes"
	"fmt"

	"github.com/martindrlik/labs/key"
)

func ExampleKeyStore() {
	const k = "k"
	secret := []byte("secret")
	s := new(key.Store)
	if s.Add(map[string]key.Value{k: {Active: true, Secret: secret}}) {
		fmt.Println("k is added")
	}
	if s.Value(k).Active {
		fmt.Println("k is active")
	}
	if bytes.Equal(s.Value(k).Secret, secret) {
		fmt.Println("k secret is secret")
	}
	if s.Deactivate(k) {
		fmt.Println("k is deactivated")
	}
	// Output:
	// k is added
	// k is active
	// k secret is secret
	// k is deactivated
}
