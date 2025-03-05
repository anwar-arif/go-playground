package car_rental_system

type PaypalPaymentProcessor struct {
}

func NewPaypalPaymentProcessor() *PaypalPaymentProcessor {
	return &PaypalPaymentProcessor{}
}

func (ppp *PaypalPaymentProcessor) ProcessPayment(amount float64) bool {
	return true
}
