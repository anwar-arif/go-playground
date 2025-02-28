package atm

import (
	"errors"
	"sync"
)

type BankingService struct {
	Accounts map[string]*Account
	mu       sync.RWMutex
}

func NewBankingService() *BankingService {
	return &BankingService{
		Accounts: make(map[string]*Account),
	}
}

func (bs *BankingService) ProcessTransaction(transaction Transaction) error {
	if err := transaction.Execute(); err != nil {
		return err
	}
	return nil
}

func (bs *BankingService) CreateAccount(accountNo string, initialBalance float64) (*Account, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	if _, ok := bs.Accounts[accountNo]; ok {
		return nil, errors.New("account already exists")
	}
	account := &Account{
		AccountNumber: accountNo,
		Balance:       initialBalance,
	}
	bs.Accounts[accountNo] = account
	return account, nil
}

func (bs *BankingService) GetAccount(accountNo string) (*Account, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	if account, ok := bs.Accounts[accountNo]; ok {
		return account, nil
	}
	return nil, errors.New("account does not exist")
}
