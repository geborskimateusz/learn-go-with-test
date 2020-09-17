package pointerserrors

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds throws an errow when withdraw is higher than balance
var ErrInsufficientFunds = errors.New("Cannot withdraw such amount")

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
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns current amount of bitcoins
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
