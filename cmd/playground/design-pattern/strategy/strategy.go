package strategy

import "fmt"

func Run() {
	context := &PaymentContext{}
	// Use credit card payment
	creditCard := NewCreditCardPayment("1234-5678-9101-1121", "123", "John Doe")
	context.SetStrategy(creditCard)

	// Try payment with retry
	result, err := context.ExecutePayment(100.50, 3)
	if err != nil {
		fmt.Println("Final payment error:", err)
	} else {
		fmt.Println(result)

		// Store transaction ID for demo purposes
		var transactionID string
		fmt.Sscanf(result, "Paid %f using Credit Card %s. Transaction ID: %s",
			new(float64), new(string), &transactionID)

		// Rollback the transaction
		rollbackResult, err := creditCard.Rollback(transactionID)
		if err != nil {
			fmt.Println("Rollback error:", err)
		} else {
			fmt.Println(rollbackResult)
		}
	}

	// Switch to PayPal with failed payment scenario
	paypal := NewPayPalPayment("john@example.com", "********")
	context.SetStrategy(paypal)

	// Try a payment that will fail (over the limit)
	result, err = context.ExecutePayment(600, 2)
	if err != nil {
		fmt.Println("Final payment error:", err)

		// Try with a valid amount
		result, err = context.ExecutePayment(400, 1)
		if err == nil {
			fmt.Println(result)
		}
	}

	// Switch to Crypto
	crypto := NewCryptoPayment("0x1234567890abcdef")
	context.SetStrategy(crypto)
	result, err = context.ExecutePayment(25.25, 1)
	if err == nil {
		fmt.Println(result)
	}
}
