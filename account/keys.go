package account

import "sync"

type Keys struct {
	sync.Mutex
	keys map[string]bool
}

// Add adds key. Key can be user name plus her password hash.
func (ks *Keys) Add(keys ...string) {
	ks.Lock()
	defer ks.Unlock()
	newSize := ks.countValid() + len(keys)
	newKeys := make(map[string]bool, newSize)
	for k, valid := range ks.keys {
		if !valid {
			continue
		}
		newKeys[k] = true
	}
	for _, k := range keys {
		newKeys[k] = true
	}
	ks.keys = newKeys
}

// IsValid returns true for valid key and false for invalid key.
func (ks *Keys) IsValid(key string) bool { return ks.keys[key] }

// Invalidate makes existing key invalid and returns true.
// It returns false if key does not exist or is invalid.
func (ks *Keys) Invalidate(key string) bool {
	if valid := ks.keys[key]; valid {
		ks.keys[key] = false
		return true
	}
	return false
}

func (ks *Keys) countValid() (count int) {
	for _, valid := range ks.keys {
		if valid {
			count++
		}
	}
	return
}
