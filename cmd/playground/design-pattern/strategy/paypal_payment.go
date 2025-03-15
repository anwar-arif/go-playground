package strategy

import (
	"errors"
	"fmt"
	"time"
)

type PayPalPayment struct {
	Email        string
	Password     string
	Transactions map[string]*Transaction
}

func NewPayPalPayment(email, password string) *PayPalPayment {
	return &PayPalPayment{
		Email:        email,
		Password:     password,
		Transactions: make(map[string]*Transaction),
	}
}

func (ppp *PayPalPayment) Pay(amount float64) (string, error) {
	// Simulate potential failure
	if amount > 500 {
		return "", errors.New("amount exceeds PayPal single transaction limit")
	}

	// Create transaction
	transactionID := fmt.Sprintf("PP-%d", time.Now().UnixNano())
	ppp.Transactions[transactionID] = &Transaction{
		TransactionId: transactionID,
		Amount:        amount,
		Timestamp:     time.Now(),
		Status:        Completed,
	}

	return fmt.Sprintf("Paid %.2f using PayPal account %s. Transaction ID: %s",
		amount, ppp.Email, transactionID), nil
}

func (ppp *PayPalPayment) Rollback(transactionID string) (string, error) {
	transaction, exists := ppp.Transactions[transactionID]
	if !exists {
		return "", errors.New("transaction not found")
	}

	if transaction.Status == RolledBack {
		return "", errors.New("transaction already rolled back")
	}

	transaction.Status = RolledBack
	return fmt.Sprintf("Rolled back transaction %s for %.2f on PayPal account %s",
		transactionID, transaction.Amount, ppp.Email), nil
}
