package match

type wantPlay struct {
	player string
	tok    chan Token
}

type server struct {
	playing  map[Token][2]string
	wantPlay chan wantPlay
}

func NewServer() *server {
	return new(server)
}

func (s *server) Listen() (cancel func()) {
	if s.playing != nil {
		panic("already listening")
	}
	s.playing = make(map[Token][2]string)
	s.wantPlay = make(chan wantPlay)
	done := make(chan struct{})
	go s.work(done)
	return func() { close(done) }
}

func (s *server) work(done <-chan struct{}) {
	c := make(chan struct{})
	go s.match(c)
	for {
		select {
		case <-c:
		case <-done:
			return
		}
	}
}

func (s *server) match(c chan struct{}) {
	for {
		a := <-s.wantPlay
		b := <-s.wantPlay
		tok := simpleToken()
		s.playing[tok] = [2]string{
			a.player,
			b.player,
		}
		a.tok <- tok
		b.tok <- tok
		close(a.tok)
		close(b.tok)
		c <- struct{}{}
	}
}

func (s *server) WantPlay(player string) <-chan Token {
	c := make(chan Token)
	go func() { s.wantPlay <- wantPlay{player, c} }()
	return c
}
