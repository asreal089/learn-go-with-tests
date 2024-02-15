package pointerserrors

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10.0)

	get := wallet.getBalance()

	want := 10.0

	if get != want {
		t.Errorf("got %.2f and want %.2f", get, want)
	}

}
