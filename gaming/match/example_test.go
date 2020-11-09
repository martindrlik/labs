package match_test

import (
	"fmt"

	"github.com/martindrlik/labs/gaming/match"
)

func ExampleServer() {
	srv := match.NewServer()
	cancel := srv.Listen()
	defer cancel()
	c := make(chan match.Token)
	go func() { c <- <-srv.WantPlay("player1") }() // player1's want play request
	go func() { c <- <-srv.WantPlay("player2") }() // player2's want play request
	fmt.Println("Players received the same match token:", <-c == <-c)
	// Output:
	// Players received the same match token: true
}
