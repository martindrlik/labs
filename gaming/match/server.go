package match

type wantPlay struct {
	player string
	ch     chan Communication
}

type server struct {
	bufferSize int
	playing    map[Token][2]string
	wantPlay   chan wantPlay
}

func NewServer() *server {
	return new(server)
}

func (s *server) Listen() (cancel func()) {
	if s.playing != nil {
		panic("already listening")
	}
	s.playing = make(map[Token][2]string)
	if s.bufferSize > 0 {
		s.wantPlay = make(chan wantPlay, s.bufferSize)
	} else {
		s.wantPlay = make(chan wantPlay)
	}
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
		a2b := make(chan string)
		b2a := make(chan string)
		a.ch <- Communication{
			MatchId: tok,
			Receive: b2a,
			Send:    a2b,
		}
		b.ch <- Communication{
			MatchId: tok,
			Receive: a2b,
			Send:    b2a,
		}
		close(a.ch)
		close(b.ch)
		c <- struct{}{}
	}
}

func (s *server) Match(player string) <-chan Communication {
	c := make(chan Communication)
	s.wantPlay <- wantPlay{player, c}
	return c
}
