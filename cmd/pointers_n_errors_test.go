package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertEqualFloatResp := func(t *testing.T, got, want CryptoCoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	assertErrorResp := func(t *testing.T, err, wantedErrMsg error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error to be return but was nil")
		}

		if err.Error() != wantedErrMsg.Error() {
			t.Errorf("wanted error msg %s, got %s", wantedErrMsg, err.Error())
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)

		got := wallet.Balance()
		want := CryptoCoin(10.00)

		assertEqualFloatResp(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)

		err := wallet.Withdraw(CryptoCoin(5.0))
		got := wallet.Balance()
		want := CryptoCoin(5.00)

		assertNoError(t, err)
		assertEqualFloatResp(t, got, want)
	})

	t.Run("withdraw error", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)

		err := wallet.Withdraw(CryptoCoin(15.0))
		assertErrorResp(t, err, WithdrawErrorMessage)
	})

	t.Run("pay", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)

		err := wallet.Pay(CryptoCoin(2.0))

		got := wallet.Balance()
		want := CryptoCoin(8.00)

		assertNoError(t, err)
		assertEqualFloatResp(t, got, want)
	})

	t.Run("pay error", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)

		err := wallet.Pay(CryptoCoin(20.0))

		assertErrorResp(t, err, PayErrorMessage)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
