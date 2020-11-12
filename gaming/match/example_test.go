package match_test

import (
	"fmt"

	"github.com/martindrlik/labs/gaming/match"
)

func ExampleServer() {
	srv := match.NewServer()
	cancel := srv.Listen()
	defer cancel()
	c := make(chan match.Communication)
	go func() { c <- <-srv.Match("player1") }() // player1 wants play
	go func() { c <- <-srv.Match("player2") }() // player2 wants play
	p1 := <-c
	p2 := <-c
	fmt.Println("Players received the same reconnect token:", p1.MatchId == p2.MatchId)
	go func() { p1.Send <- "hello" }()
	fmt.Printf("Player1 sent %q to player2\n", <-p2.Receive)
	// Output:
	// Players received the same reconnect token: true
	// Player1 sent "hello" to player2
}
