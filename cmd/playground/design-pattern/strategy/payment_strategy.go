package strategy

type PaymentStrategy interface {
	Pay(amount float64) (string, error)
	Rollback(transactionId string) (string, error)
}
