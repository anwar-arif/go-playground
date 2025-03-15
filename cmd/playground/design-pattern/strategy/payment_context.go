package strategy

import (
	"errors"
	"fmt"
	"time"
)

type PaymentContext struct {
	strategy          PaymentStrategy
	lastTransactionId string
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64, maxRetries int) (string, error) {
	var result string
	var err error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			fmt.Printf("Retrying payment (attempt %d of %d)...\n", attempt, maxRetries)
			time.Sleep(time.Duration(attempt*500) * time.Millisecond)
		}

		if result, err = pc.strategy.Pay(amount); err == nil {
			return result, err
		}
		fmt.Printf("Payment attempt failed: %v\n", err.Error())
	}
	return "", fmt.Errorf("payment failed after %d attempts: %w", maxRetries+1, err)
}

// RollbackLastPayment rolls back the last successful transaction
func (pc *PaymentContext) RollbackLastPayment() (string, error) {
	if pc.lastTransactionId == "" {
		return "", errors.New("no transaction to roll back")
	}

	return pc.strategy.Rollback(pc.lastTransactionId)
}
