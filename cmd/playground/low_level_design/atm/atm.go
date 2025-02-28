package atm

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

type ATM struct {
	CashDispenser  *CashDispenser
	BankingService *BankingService
	TxnCounter     int64
}

func NewATM(cashDispenser *CashDispenser, bankingService *BankingService) *ATM {
	return &ATM{
		CashDispenser:  cashDispenser,
		BankingService: bankingService,
		TxnCounter:     0,
	}
}

func (atm *ATM) AuthenticateUser(card *Card) error {
	// Authenticate logic here
	return nil
}

func (atm *ATM) CheckBalance(accountNumber string) (float64, error) {
	account, _ := atm.BankingService.GetAccount(accountNumber)
	if account == nil {
		return 0, errors.New("account not found")
	}
	return account.GetBalance(), nil
}

func (atm *ATM) GenerateTransactionId() string {
	txnNumber := atomic.AddInt64(&atm.TxnCounter, 1)
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("TXN%s%010d", timestamp, txnNumber)
}

func (atm *ATM) WithdrawCash(accountNumber string, amount float64) error {
	account, err := atm.BankingService.GetAccount(accountNumber)
	if err != nil {
		return errors.New("account not found")
	}

	transaction := NewWithdrawalTransaction(atm.GenerateTransactionId(), account, amount)
	if err := atm.BankingService.ProcessTransaction(transaction); err != nil {
		return err
	}

	return atm.CashDispenser.DispenseCash(int(amount))
}

func (atm *ATM) DepositCash(accountNumber string, amount float64) error {
	account, err := atm.BankingService.GetAccount(accountNumber)
	if err != nil {
		return errors.New("account not found")
	}

	transaction := NewDepositTransaction(atm.GenerateTransactionId(), account, amount)
	return atm.BankingService.ProcessTransaction(transaction)
}
