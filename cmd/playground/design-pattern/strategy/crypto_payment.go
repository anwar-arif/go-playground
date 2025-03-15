package strategy

import (
	"errors"
	"fmt"
	"time"
)

type CryptoPayment struct {
	WalletId     string
	Transactions map[string]*Transaction
}

func NewCryptoPayment(walletId string) *CryptoPayment {
	return &CryptoPayment{
		WalletId:     walletId,
		Transactions: make(map[string]*Transaction),
	}
}

func (cp *CryptoPayment) Pay(amount float64) (string, error) {
	// Simulate potential failure
	if amount > 10000 {
		return "", errors.New("amount requires additional verification for crypto transaction")
	}

	// Create transaction
	transactionID := fmt.Sprintf("CR-%d", time.Now().UnixNano())
	cp.Transactions[transactionID] = &Transaction{
		TransactionId: transactionID,
		Amount:        amount,
		Timestamp:     time.Now(),
		Status:        Completed,
	}

	return fmt.Sprintf("Paid %.2f using Crypto wallet %s. Transaction ID: %s",
		amount, cp.WalletId, transactionID), nil
}

func (cp *CryptoPayment) Rollback(transactionId string) (string, error) {
	transaction, exists := cp.Transactions[transactionId]
	if !exists {
		return "", errors.New("transaction not found")
	}

	if transaction.Status == RolledBack {
		return "", errors.New("transaction is already rolled back")
	}

	transaction.Status = RolledBack
	return fmt.Sprintf("Rolled back transaction %s for %.2f on Crypto wallet %s", transactionId, transaction.Amount, cp.WalletId), nil
}
