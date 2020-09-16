package test

import (
	"testing"

	pointerserrors "geborskimateusz.com/m/pointerserrors"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet pointerserrors.Wallet, want pointerserrors.Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit bitcoins", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}
		wallet.Deposit(pointerserrors.Bitcoin(10))
		want := pointerserrors.Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw bitcoins", func(t *testing.T) {
		wallet := pointerserrors.Wallet{balance: pointerserrors.Bitcoin(10)}
		wallet.Withdraw(pointerserrors.Bitcoin(5))
		want := pointerserrors.Bitcoin(5)

		assertBalance(t, wallet, want)
	})

}
