package samples

import "sync"

type Account struct {
	mu      sync.Mutex
	Balance int64
}

func (a *Account) Deposit(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance -= amount
}
