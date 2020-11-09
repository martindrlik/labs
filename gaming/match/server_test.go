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
	c := make(chan match.Token)
	go func() { c <- <-s.WantPlay("player1") }()
	go func() { c <- <-s.WantPlay("player2") }()
	t1 := (<-c).String()
	t2 := (<-c).String()
	if t1 == "" || t2 == "" {
		t.Errorf("expected players to get non empty token: t1(%q) t2(%q)", t1, t2)
	}
	if t1 == "" || t1 != t2 {
		t.Errorf("expected players to get the same token: t1(%q) t2(%q)", t1, t2)
	}
}
