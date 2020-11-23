package account_test

import (
	"testing"
	"time"

	"github.com/martindrlik/labs/account"
)

func TestSessions(t *testing.T) {
	ss := &account.Sessions{
		Now: time.Now,
	}
	const sessionKey = "sessionKey"
	t.Run("No session can be valid", func(t *testing.T) {
		if ss.IsValid(sessionKey) {
			t.Error("There is no session so no session can be valid")
		}
	})
	t.Run("Add", func(t *testing.T) {
		ss.Add(map[string]account.Session{sessionKey: {
			Active:  true,
			ValidTo: time.Now().Add(time.Minute),
		}})
	})
	t.Run("Session is valid", func(t *testing.T) {
		if !ss.IsValid(sessionKey) {
			t.Error("Session should be valid")
		}
	})
	t.Run("Deactivate true", func(t *testing.T) {
		if !ss.Deactivate(sessionKey) {
			t.Error("Session key should be invalidated")
		}
	})
	t.Run("Deactivate false", func(t *testing.T) {
		if ss.Deactivate(sessionKey) {
			t.Error("Session key cannot be invalidated as it should be already deactivated")
		}
	})
}
