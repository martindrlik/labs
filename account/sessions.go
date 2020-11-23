package account

import (
	"sync"
	"time"
)

type Session struct {
	Active  bool
	ValidTo time.Time
}

func (s Session) IsValid(time time.Time) bool {
	return s.Active && time.Before(s.ValidTo)
}

type Sessions struct {
	sync.Mutex
	online map[string]Session

	Now func() time.Time
}

// Add adds session.
func (ss *Sessions) Add(sessions map[string]Session) {
	ss.Lock()
	defer ss.Unlock()
	now := ss.Now()
	newSize := sessionSize(now, ss.online) + len(sessions)
	newSessions := make(map[string]Session, newSize)
	for k, v := range ss.online {
		if !v.IsValid(now) {
			continue
		}
		newSessions[k] = v
	}
	for k, v := range sessions {
		newSessions[k] = v
	}
	ss.online = newSessions
}

// IsValid returns true for valid and false for invalid session key.
func (ss *Sessions) IsValid(key string) bool {
	return ss.online[key].IsValid(ss.Now())
}

// Deactivate deactivates active session and returns true.
// It returns false if session does not exist or is already deactivated.
func (ss *Sessions) Deactivate(key string) bool {
	s, ok := ss.online[key]
	if !ok || !s.IsValid(ss.Now()) {
		return false
	}
	s.Active = false
	ss.online[key] = s
	return true
}

func sessionSize(time time.Time, m map[string]Session) (size int) {
	for _, v := range m {
		if !v.IsValid(time) {
			continue
		}
		size++
	}
	return
}
