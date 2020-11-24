package session_test

import (
	"fmt"
	"time"

	"github.com/martindrlik/labs/session"
)

func ExampleSessionStore() {
	const a = "a"
	now := time.Now()
	validTo := now.Add(time.Minute)
	s := &session.Store{Now: time.Now}
	if s.Add(map[string]session.Value{a: {Active: true, ValidTo: validTo}}) {
		fmt.Println("a is added")
	}
	if s.Value(a).IsValid(now) {
		fmt.Println("a is valid")
	}
	if s.Value(a).ValidTo.Equal(validTo) {
		fmt.Println("a's validTo is validTo")
	}
	if s.Deactivate(a) {
		fmt.Println("a is deactivated")
	}
	// Output:
	// a is added
	// a is valid
	// a's validTo is validTo
	// a is deactivated
}
