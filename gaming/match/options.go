package match

type option func(*server) option

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (s *server) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(s)
	}
	return previous
}

// BufferSize sets server's want play channel buffer size to n.
func BufferSize(n int) option {
	return func(s *server) option {
		previous := s.bufferSize
		s.bufferSize = n
		return BufferSize(previous)
	}
}
