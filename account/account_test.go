package account_test

import (
	"strings"
	"testing"

	"github.com/martindrlik/labs/account"
)

func TestAccounts(t *testing.T) {
	const a = "a"
	s := new(account.Store)
	t.Run("Add(a)(true)", func(t *testing.T) {
		if !s.Add(map[string]account.Account{a: {Active: true}}) {
			t.Fatal("should be possible to add")
		}
	})
	t.Run("IsAvailable(a)(false)", func(t *testing.T) {
		if s.IsAvailable(a) {
			t.Error("a should not be available")
		}
	})
	t.Run("Add(a)(false)", func(t *testing.T) {
		if s.Add(map[string]account.Account{a: {Active: true}}) {
			t.Error("already added so it should not be possible to add")
		}
	})
	t.Run("Account(a).Active", func(t *testing.T) {
		if !s.Account(a).Active {
			t.Error("should be active")
		}
	})
	t.Run("Deactivate(a)(true)", func(t *testing.T) {
		if !s.Deactivate(a) {
			t.Error("should be possible to deactivate")
		}
	})
	t.Run("Deactivate(a)(false)", func(t *testing.T) {
		if s.Deactivate(a) {
			t.Error("already deactivated so it cannot be deactivated")
		}
	})
	t.Run("Deactivate(a)(false)", func(t *testing.T) {
		if s.Deactivate(a) {
			t.Error("already deactivated so it cannot be deactivated")
		}
	})
}

func TestNames(t *testing.T) {
	s := new(account.Store)
	if !s.Add(map[string]account.Account{
		"e": {Active: true},
		"d": {},
		"c": {Active: true},
		"b": {Active: true},
		"a": {Active: true},
	}) {
		t.Fatal("unable to add")
	}
	a := strings.Join(s.Names(), "")
	if a != "abcde" {
		t.Errorf("expected sorted names abcde, got %v", a)
	}
}
