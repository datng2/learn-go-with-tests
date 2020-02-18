package wallet

import "errors"

var (
	// ErrInsufficientFunds an account
	ErrInsufficientFunds = errors.New("Overdrafted")
)

// Wallet is an abstraction of personal banking account
type Wallet struct {
	// In Go if a symbol (so variables, types, functions et al) starts with
	//a lowercase symbol then it is private outside the package it's defined in.

	// private outside the package
	balance Bitcoin

	// public outside the package
	SupportedCurrencies []int
}

// Deposit updates balance account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

// Withdraw gets latest balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance gets latest balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
