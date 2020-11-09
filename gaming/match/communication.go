package match

type Communication struct {
	Recn Token
	Recv <-chan string
	Send chan<- string
}
