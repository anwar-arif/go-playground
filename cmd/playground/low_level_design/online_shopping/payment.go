package online_shopping

type Payment interface {
	ProcessPayment(amount float64) bool
}
