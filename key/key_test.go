package key_test

import (
	"strings"
	"testing"

	"github.com/martindrlik/labs/key"
)

func TestKeys(t *testing.T) {
	const k = "k"
	s := new(key.Store)
	t.Run("Add(k)(true)", func(t *testing.T) {
		if !s.Add(map[string]key.Value{k: {Active: true}}) {
			t.Fatal("should be possible to add")
		}
	})
	t.Run("Add(k)(false)", func(t *testing.T) {
		if s.Add(map[string]key.Value{k: {Active: true}}) {
			t.Error("already added so it should not be possible to add")
		}
	})
	t.Run("Value(k).Active", func(t *testing.T) {
		if !s.Value(k).Active {
			t.Error("should be active")
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
	s := new(key.Store)
	if !s.Add(map[string]key.Value{
		"e": {Active: true},
		"d": {},
		"c": {Active: true},
		"b": {Active: true},
		"a": {Active: true},
	}) {
		t.Fatal("unable to add")
	}
	a := strings.Join(s.Names(), "")
	if a != "abce" {
		t.Errorf("expected sorted active names abce, got %v", a)
	}
}
