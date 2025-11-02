package transactions

import (
	"sort"
	"time"
)

type Transaction struct {
	UserId    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	TimeStamp int64   `json:"timestamp"`
}

type Validator interface {
	Validate(transaction Transaction) bool
}

type FraudulentValidator struct {
	Threshold        float64
	Window           time.Duration
	UserTransactions map[string][]Transaction
}

func (fv FraudulentValidator) Validate(allTransactions []Transaction) []string {
	sort.Slice(allTransactions, func(i, j int) bool {
		return allTransactions[i].TimeStamp < allTransactions[j].TimeStamp
	})
	for _, transaction := range allTransactions {
		userId := transaction.UserId
		fv.UserTransactions[userId] = append(fv.UserTransactions[userId], transaction)
	}

	users := make([]string, 0)
	for userId, transactions := range fv.UserTransactions {
		n := len(transactions)
		i, j := 0, -1
		amount := 0.0
		for ; i < n; i = i + 1 {
			for j+1 < n && time.Duration(transactions[j+1].TimeStamp-transactions[i].TimeStamp) <= fv.Window {
				j += 1
				amount += transactions[j].Amount
			}
			if amount > fv.Threshold {
				users = append(users, userId)
				break
			}
			amount -= transactions[i].Amount
		}
	}

	return users
}

func DetectFraudulentUsers(transactions []Transaction, threshold float64, window time.Duration) []string {
	fv := FraudulentValidator{
		Threshold:        threshold,
		Window:           window,
		UserTransactions: make(map[string][]Transaction),
	}
	return fv.Validate(transactions)
}
