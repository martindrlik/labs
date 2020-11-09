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
	go func() { c <- <-srv.WantPlay("player1") }() // player1's want play request
	go func() { c <- <-srv.WantPlay("player2") }() // player2's want play request
	p1 := <-c
	p2 := <-c
	fmt.Println("Players received the same reconnect token:", p1.Recn == p2.Recn)
	go func() { p1.Send <- "hello" }()
	fmt.Printf("Player1 sent %q to player2\n", <-p2.Recv)
	// Output:
	// Players received the same reconnect token: true
	// Player1 sent "hello" to player2
}
