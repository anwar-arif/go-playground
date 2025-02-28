package atm

import (
	"errors"
	"sync"
)

type Account struct {
	AccountNumber string
	Balance       float64
	mu            sync.RWMutex
}

func NewAccount(number string, balance float64) *Account {
	return &Account{
		AccountNumber: number,
		Balance:       balance,
	}
}

func (a *Account) GetAccountNumber() string {
	return a.AccountNumber
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	balance := a.Balance
	return balance
}

func (a *Account) Debit(balance float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance < balance {
		return errors.New("not enough balance")
	}
	a.Balance -= balance
	return nil
}

func (a *Account) Credit(balance float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += balance
	return nil
}
