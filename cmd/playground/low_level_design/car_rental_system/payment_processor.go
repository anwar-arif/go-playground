package car_rental_system

type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}
