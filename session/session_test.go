package session_test

import (
	"strings"
	"testing"
	"time"

	"github.com/martindrlik/labs/session"
)

func TestStore(t *testing.T) {
	const k = "k"
	now := time.Now()
	validTo := now.Add(time.Minute)
	s := &session.Store{Now: time.Now}
	t.Run("Add(k)(true)", func(t *testing.T) {
		if !s.Add(map[string]session.Value{k: {Active: true, ValidTo: validTo}}) {
			t.Fatal("should be possible to add")
		}
	})
	t.Run("Add(k)(false)", func(t *testing.T) {
		if s.Add(map[string]session.Value{k: {Active: true}}) {
			t.Error("already added so it should not be possible to add")
		}
	})
	t.Run("Value(k).IsValid(now)", func(t *testing.T) {
		if !s.Value(k).IsValid(now) {
			t.Error("should be valid")
		}
	})
	t.Run("Deactivate(k)(true)", func(t *testing.T) {
		if !s.Deactivate(k) {
			t.Error("should be possible to deactivate")
		}
	})
	t.Run("Deactivate(k)(false)", func(t *testing.T) {
		if s.Deactivate(k) {
			t.Error("already deactivated so it cannot be deactivated")
		}
	})
	t.Run("Deactivate(k)(false)", func(t *testing.T) {
		if s.Deactivate(k) {
			t.Error("already deactivated so it cannot be deactivated")
		}
	})
}

func TestNames(t *testing.T) {
	now := time.Now()
	validTo := now.Add(time.Minute)
	s := &session.Store{Now: func() time.Time { return now }}
	if !s.Add(map[string]session.Value{
		"e": {Active: true, ValidTo: validTo},
		"d": {ValidTo: validTo},
		"c": {Active: true, ValidTo: validTo},
		"b": {Active: true, ValidTo: validTo},
		"a": {Active: true},
	}) {
		t.Fatal("unable to add")
	}
	a := strings.Join(s.Names(), "")
	if a != "bce" {
		t.Errorf("expected sorted active names bce, got %v", a)
	}
}
