package online_shopping

type CreditCardPayment struct {
}

func NewCreditCardPayment() *CreditCardPayment {
	return &CreditCardPayment{}
}

func (ccp *CreditCardPayment) ProcessPayment(amount float64) bool {
	return true
}
