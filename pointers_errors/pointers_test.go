package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.getBalance()
		if got != want {
			t.Errorf("got %s and want %s", wallet.getBalance(), want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		if err != nil {
			assertBalance(t, wallet, Bitcoin(20))
			assertError(t, err, "cannot withdraw, insufficient funds")
		} else {
			t.Error("wanted an error but didn't get one")
		}
	})
}
