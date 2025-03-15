package strategy

import (
	"errors"
	"fmt"
	"time"
)

type CreditCardPayment struct {
	CardNumber   string
	CVV          string
	Name         string
	Transactions map[string]*Transaction
}

func NewCreditCardPayment(cardNumber, cvv, name string) *CreditCardPayment {
	return &CreditCardPayment{
		CardNumber:   cardNumber,
		CVV:          cvv,
		Name:         name,
		Transactions: make(map[string]*Transaction),
	}
}

func (ccp *CreditCardPayment) Pay(amount float64) (string, error) {
	if amount > 1000 {
		return "", errors.New("amount exceeds credit card limit")
	}

	transactionId := fmt.Sprintf("cc-%d", time.Now().UnixNano())
	ccp.Transactions[transactionId] = &Transaction{
		TransactionId: transactionId,
		Amount:        amount,
		Timestamp:     time.Now(),
		Status:        Completed,
	}
	return fmt.Sprintf("Paid %.2f using credit card %s. Transaction Id: %s", amount, ccp.CardNumber, transactionId), nil
}

func (ccp *CreditCardPayment) Rollback(transactionId string) (string, error) {
	transaction, exists := ccp.Transactions[transactionId]
	if !exists {
		return "", errors.New("transaction not found")
	}

	if transaction.Status == RolledBack {
		return "", errors.New("transaction is already rolled back")
	}

	transaction.Status = RolledBack
	return fmt.Sprintf("Rolled back transaction %s for %.2f on Credit Card %s",
		transactionId, transaction.Amount, ccp.CardNumber), nil
}
