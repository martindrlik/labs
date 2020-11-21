package player_test

import (
	"fmt"

	"github.com/martindrlik/labs/gaming/player"
)

func ExamplePlayers() {
	ps := new(player.Players)
	ps.Register(map[string]chan string{"Amanda": make(chan string, 1)})
	amanda, ok := ps.Player("Amanda")
	if !ok {
		fmt.Println("Amanda is not registered!")
		return
	}
	done := make(chan struct{})
	go func() {
		fmt.Printf("Amanda received %s\n", <-amanda)
		close(done)
	}()
	fmt.Println("Sending hello to Amanda")
	amanda <- "hello"
	<-done
	// Output:
	// Sending hello to Amanda
	// Amanda received hello
}
