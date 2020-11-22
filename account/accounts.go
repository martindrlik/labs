package account

import "sync"

type Account struct {
	Inactive bool
}

type Accounts struct {
	sync.Mutex
	byName map[string]Account
}

// Add adds an account.
func (as *Accounts) Add(accounts map[string]Account) {
	as.Lock()
	defer as.Unlock()
	newSize := len(as.byName) + len(accounts)
	newAccounts := make(map[string]Account, newSize)
	for k, v := range as.byName {
		newAccounts[k] = v
	}
	for k, v := range accounts {
		newAccounts[k] = v
	}
	as.byName = newAccounts
}

// IsAvailable returns true if given name is available.
// It returns false if given name is already used.
func (as *Accounts) IsAvailable(name string) bool {
	_, used := as.byName[name]
	return !used
}

// Deactivate deactivates active account and returns true.
// It returns false if account does not exist or is inactive.
func (as *Accounts) Deactivate(name string) bool {
	a, ok := as.byName[name]
	if !ok || a.Inactive {
		return false
	}
	a.Inactive = true
	as.byName[name] = a
	return true
}
