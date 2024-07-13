package main

import (
	"errors"
	"fmt"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

type CryptoCoin float64

var PayErrorMessage = errors.New("cannot pay specified amount")
var WithdrawErrorMessage = errors.New("cannot withdraw specified amount")

type Wallet struct {
	money CryptoCoin
}

func (w *Wallet) Pay(amount CryptoCoin) error {
	if amount > w.money {
		return PayErrorMessage
	}
	w.money -= amount
	return nil
}

func (w *Wallet) Deposit(amount CryptoCoin) {
	w.money += amount
}

func (w *Wallet) Balance() CryptoCoin {
	return w.money
}

func (w *Wallet) Withdraw(amount CryptoCoin) error {
	if w.money < amount {
		return WithdrawErrorMessage
	}
	w.money -= amount
	return nil
}

// implementing the Stringer interface defined in go.
// this lets you determine how your objects/types are printed
func (w *Wallet) String() string {
	return fmt.Sprintf("my Crypto BAL: %.2f", w.money)
}
