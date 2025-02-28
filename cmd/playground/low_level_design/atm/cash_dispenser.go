package atm

import (
	"errors"
	"sync"
)

type CashDispenser struct {
	CashAvailable int
	mu            sync.RWMutex
}

func NewCashDispenser(cash int) *CashDispenser {
	return &CashDispenser{
		CashAvailable: cash,
	}
}

func (cd *CashDispenser) DispenseCash(amount int) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	if amount > cd.CashAvailable {
		return errors.New("insufficient cash in ATM")
	}
	cd.CashAvailable -= amount
	return nil
}
