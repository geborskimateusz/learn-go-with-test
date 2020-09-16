package test

import (
	"testing"

	pointerserrors "geborskimateusz.com/m/pointerserrors"
)

func TestWallet(t *testing.T) {

	wallet := pointerserrors.Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
