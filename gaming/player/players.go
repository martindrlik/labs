package player

type Players struct {
	online map[string]chan string
}

func (ps *Players) Register(players map[string]chan string) {
	newSize := len(ps.online) + len(players)
	newData := make(map[string]chan string, newSize)
	for k, v := range ps.online {
		newData[k] = v
	}
	for k, v := range players {
		newData[k] = v
	}
	ps.online = newData
}

func (ps *Players) Player(name string) (ch chan string, ok bool) {
	ch, ok = ps.online[name]
	return
}
