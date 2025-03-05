package car_rental_system

type CreditCardPaymentProcessor struct {
}

func NewCreditCardPaymentProcessor() *CreditCardPaymentProcessor {
	return &CreditCardPaymentProcessor{}
}

func (ccpp *CreditCardPaymentProcessor) ProcessPayment(amount float64) bool {
	return true
}
