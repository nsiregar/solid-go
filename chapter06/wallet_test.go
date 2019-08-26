package wallet_test

import (
	"testing"

	"github.com/nsiregar/solid-go/chapter06"
)

func TestChapter06Wallet(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected wallet.Bitcoin) {
		t.Helper()

		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		ewallet := wallet.Wallet{}
		amount := wallet.Bitcoin(10)

		ewallet.Deposit(amount)

		result := ewallet.Balance()
		expected := amount

		assertEqual(t, result, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		amount := wallet.Bitcoin(100)
		ewallet := wallet.Wallet{}

		ewallet.Deposit(amount)
		ewallet.Withdraw(wallet.Bitcoin(10))

		result := ewallet.Balance()
		expected := wallet.Bitcoin(90)

		assertEqual(t, result, expected)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		amount := wallet.Bitcoin(100)
		ewallet := wallet.Wallet{}

		ewallet.Deposit(amount)
		err := ewallet.Withdraw(wallet.Bitcoin(1000))

		assertEqual(t, ewallet.Balance(), amount)

		if err == nil {
			t.Error("should raise error, but do not get any")
		}
	})
}
