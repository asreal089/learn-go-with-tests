package pointerserrors

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) getBalance() float64 {
	return w.balance
}
