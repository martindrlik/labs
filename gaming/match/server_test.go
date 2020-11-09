package match_test

import (
	"testing"

	"github.com/martindrlik/labs/gaming/match"
)

func TestCannotListenTwice(t *testing.T) {
	defer func() {
		x, ok := recover().(string)
		if !ok || x != "already listening" {
			t.Fatal("expected panic(\"already listening\") for multiple Listen calls")
		}
	}()
	s := match.NewServer()
	defer s.Listen()()
	defer s.Listen()()
}

func TestWantPlay(t *testing.T) {
	s := match.NewServer()
	defer s.Listen()()
	c := make(chan match.Communication)
	go func() { c <- <-s.WantPlay("player1") }()
	go func() { c <- <-s.WantPlay("player2") }()
	p1 := <-c
	p2 := <-c
	t1 := p1.Recn.String()
	t2 := p2.Recn.String()
	t.Run("non empty reconnect token", func(t *testing.T) {
		if t1 == "" || t2 == "" {
			t.Errorf("expected players to get non empty token: t1(%q) t2(%q)", t1, t2)
		}
	})
	t.Run("same reconnect token", func(t *testing.T) {
		if t1 == "" || t1 != t2 {
			t.Errorf("expected players to get the same token: t1(%q) t2(%q)", t1, t2)
		}
	})
	t.Run("send", func(t *testing.T) {
		go func() { p1.Send <- "hello" }()
	})
	t.Run("retrieve", func(t *testing.T) {
		msg := <-p2.Recv
		if msg != "hello" {
			t.Errorf("expected to receive \"hello\", got %s", msg)
		}
	})
}
