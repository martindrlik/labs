package account

import (
	"sort"
	"sync"
)

type Account struct {
	Active bool
	Salt   []byte
}

type Store struct {
	sync.Mutex
	data map[string]Account
}

// Add adds accounts and returns true if all added.
// It adds nothing and returns false if some account already exists.
func (s *Store) Add(accounts map[string]Account) bool {
	s.Lock()
	defer s.Unlock()
	l := len(s.data) + len(accounts)
	d := make(map[string]Account, l)
	for k, v := range s.data {
		d[k] = v
	}
	for k, v := range accounts {
		_, ok := d[k]
		if ok {
			return false
		}
		d[k] = v
	}
	s.data = d
	return true
}

// Value returns account given by its name and true.
// It returns default value if there is no account for given name.
func (s *Store) Value(name string) Account { return s.data[name] }

// IsAvailable returns true if given name is available.
// It returns false if given name is already used.
func (s *Store) IsAvailable(name string) bool {
	_, used := s.data[name]
	return !used
}

// Deactivate deactivates existing and active account and returns true.
// It returns false if there is no account for given name or it is not active.
func (s *Store) Deactivate(name string) bool {
	if v := s.data[name]; v.Active {
		v.Active = false
		s.data[name] = v
		return true
	}
	return false
}

// Names returns sorted active account names.
func (s *Store) Names() []string {
	d := s.data
	a := make([]string, 0, len(d))
	for k := range d {
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}
