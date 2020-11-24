package key

import (
	"sort"
	"sync"
)

type Value struct {
	Active bool
	Secret []byte
}

type Store struct {
	sync.Mutex
	data map[string]Value
}

// Add adds keys and returns true if all added.
// It adds nothing and returns false if some key already exists.
func (s *Store) Add(keys map[string]Value) bool {
	s.Lock()
	defer s.Unlock()
	l := activeSize(s.data) + len(keys)
	d := make(map[string]Value, l)
	for k, v := range s.data {
		if !v.Active {
			continue
		}
		d[k] = v
	}
	for k, v := range keys {
		if _, ok := d[k]; ok {
			return false
		}
		d[k] = v
	}
	s.data = d
	return true
}

func activeSize(d map[string]Value) (size int) {
	for _, v := range d {
		if v.Active {
			size++
		}
	}
	return
}

// Value returns key given by its name and true.
// It returns default value if there is no key for given name.
func (s *Store) Value(name string) Value { return s.data[name] }

// Deactivate deactivates existing and active key and returns true.
// It returns false if there is no key for given name or it is not active.
func (s *Store) Deactivate(name string) bool {
	if v := s.data[name]; v.Active {
		v.Active = false
		s.data[name] = v
		return true
	}
	return false
}

// Names returns sorted active names.
func (s *Store) Names() []string {
	d := s.data
	a := make([]string, 0, activeSize(d))
	for k, v := range d {
		if !v.Active {
			continue
		}
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}
