package atm

type Transaction interface {
	Execute() error
}

type BaseTransaction struct {
	TransactionId string
	Account       *Account
	Amount        float64
}
