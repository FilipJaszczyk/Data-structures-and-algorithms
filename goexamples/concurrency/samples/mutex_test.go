package samples

import (
	"sync"
	"testing"
)

func TestAccountConcurrent(t *testing.T) {
	acc := &Account{
		Balance: 0,
	}
	var wg sync.WaitGroup
	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			defer wg.Done()
			acc.Deposit(1)
		}
	}()
	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			defer wg.Done()
			acc.Withdraw(1)
		}
	}()
	wg.Wait()
	if acc.Balance != 0 {
		t.Fatalf("expected account balance is %d got %d ", 0, acc.Balance)
	}
}
