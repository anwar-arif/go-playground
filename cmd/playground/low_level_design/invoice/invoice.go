package invoice

import "time"

type Purchase struct {
	UserId    string        `json:"user_id"`
	Amount    float64       `json:"amount"`
	TimeStamp time.Duration `json:"time_stamp"`
}

type Invoice struct {
	UserId string  `json:"userId"`
	Total  float64 `json:"total"`
}

type InvoiceCalculator struct {
	Discount float64 `json:"discount"`
	TaxRate  float64 `json:"tax_rate"`
}

func (ic *InvoiceCalculator) Calculate(purchases []Purchase) map[string]Invoice {
	userTotals := make(map[string]Invoice, 0)
	for _, purchase := range purchases {
		userId := purchase.UserId
		invoice := userTotals[userId]
		invoice.Total += purchase.Amount
		userTotals[userId] = invoice
	}

	for userId, invoice := range userTotals {
		invoice.Total *= ic.Discount
		invoice.Total *= ic.TaxRate
		invoice.UserId = userId
		userTotals[userId] = invoice
	}
	return userTotals
}

func GenerateInvoices(purchases []Purchase, taxRate float64, discountRate float64) map[string]Invoice {
	ic := InvoiceCalculator{
		Discount: discountRate,
		TaxRate:  taxRate,
	}

	return ic.Calculate(purchases)
}
