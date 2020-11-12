package match

type Communication struct {
	MatchId Token
	Receive <-chan string
	Send    chan<- string
}
