package pointerserrors

import (
	"fmt"
)

// Bitcoin represents number of bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores number of bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit amount of bitcoins
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw amount of bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance returns current amount of bitcoins
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
