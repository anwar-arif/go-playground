package strategy

import "time"

type TransactionStatus string

const (
	Completed  TransactionStatus = "Completed"
	RolledBack TransactionStatus = "Rolled_back"
)

type Transaction struct {
	TransactionId string
	Amount        float64
	Timestamp     time.Time
	Status        TransactionStatus
}
