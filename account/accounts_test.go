package account_test

import (
	"testing"

	"github.com/martindrlik/labs/account"
)

const amanda = "amanda"

func TestAccounts(t *testing.T) {
	as := new(account.Accounts)
	t.Run("IsAvailable false", func(t *testing.T) {
		if !as.IsAvailable(amanda) {
			t.Errorf("Name %v should be available as it is not yet used", amanda)
		}
	})
	t.Run("Add", func(t *testing.T) {
		as.Add(map[string]account.Account{amanda: {Active: true}})
	})
	t.Run("IsAvailable true", func(t *testing.T) {
		if as.IsAvailable(amanda) {
			t.Errorf("Name %v should not be available as it is already used", amanda)
		}
	})
	t.Run("Deactivate true", func(t *testing.T) {
		if !as.Deactivate(amanda) {
			t.Errorf("%v's account should be possible to deactivate", amanda)
		}
	})
	t.Run("Deactivate false", func(t *testing.T) {
		if as.Deactivate(amanda) {
			t.Errorf("%v's account cannot be deativated as it is no longer active", amanda)
		}
	})
}
