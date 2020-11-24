package session

import (
	"sort"
	"sync"
	"time"
)

type Value struct {
	Active  bool
	ValidTo time.Time
}

func (s Value) IsValid(time time.Time) bool {
	return s.Active && time.Before(s.ValidTo)
}

type Store struct {
	sync.Mutex
	data map[string]Value

	Now func() time.Time
}

// Add adds sessions and returns true if all added.
// It adds nothing and returns false if some session already exists.
func (s *Store) Add(sessions map[string]Value) bool {
	s.Lock()
	defer s.Unlock()
	t := s.Now()
	l := activeSize(t, s.data) + len(sessions)
	d := make(map[string]Value, l)
	for k, v := range s.data {
		if !v.IsValid(t) {
			continue
		}
		d[k] = v
	}
	for k, v := range sessions {
		if _, ok := d[k]; ok {
			return false
		}
		d[k] = v
	}
	s.data = d
	return true
}

func activeSize(t time.Time, d map[string]Value) (size int) {
	for _, v := range d {
		if v.IsValid(t) {
			size++
		}
	}
	return
}

// Value returns session given by its name and true.
// It returns default value if there is no session for given name.
func (s *Store) Value(name string) Value { return s.data[name] }

// Deactivate deactivates existing and active session and returns true.
// It returns false if there is no session for given name or it is not active.
func (s *Store) Deactivate(name string) bool {
	now := s.Now()
	if v := s.data[name]; v.IsValid(now) {
		v.Active = false
		s.data[name] = v
		return true
	}
	return false
}

// Names returns sorted active session names.
func (s *Store) Names() []string {
	d := s.data
	t := s.Now()
	a := make([]string, 0, activeSize(t, d))
	for k, v := range d {
		if !v.IsValid(t) {
			continue
		}
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}
