package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		want := Bitcoin(5)

		assertBalance(t, wallet, want)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted error got nill")
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, err error) {

	if err != nil {
		t.Fatal("Wanted nill got error")
	}
}
