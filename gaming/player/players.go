package player

type Players struct {
	data map[string]chan string
}

func (ps *Players) Register(players map[string]chan string) {
	newSize := len(ps.data) + len(players)
	newData := make(map[string]chan string, newSize)
	for k, v := range ps.data {
		newData[k] = v
	}
	for k, v := range players {
		newData[k] = v
	}
	ps.data = newData
}

func (ps *Players) Player(name string) (ch chan string, ok bool) {
	ch, ok = ps.data[name]
	return
}
