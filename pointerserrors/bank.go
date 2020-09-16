package pointerserrors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(ammount int) {
	w.balance += ammount
}

func (w *Wallet) Balance() int {
	return w.balance
}
