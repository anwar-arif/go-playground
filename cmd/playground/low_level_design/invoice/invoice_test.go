package invoice

import (
	"testing"
)

func similar(first map[string]Invoice, second map[string]Invoice) bool {
	count := make(map[string]float64, 0)
	for _, inv := range first {
		count[inv.UserId] += inv.Total
	}

	for _, inv := range second {
		count[inv.UserId] -= inv.Total
	}

	for userId, difference := range count {
		if userId == "" {
			continue
		}
		if difference != 0.0 {
			return false
		}
	}
	return true
}

func TestGenerateInvoices(t *testing.T) {
	tests := []struct {
		purchases    []Purchase
		taxRate      float64
		discountRate float64
		expected     map[string]Invoice
	}{
		{
			purchases: []Purchase{
				{UserId: "u1", Amount: 100.0},
				{UserId: "u2", Amount: 200.0},
				{UserId: "u1", Amount: 50.0},
				{UserId: "u3", Amount: 25.0},
			},
			taxRate:      1.0,
			discountRate: 1.0,
			expected: map[string]Invoice{
				"u1": {UserId: "u1", Total: 150.0},
				"u2": {UserId: "u2", Total: 200.0},
				"u3": {UserId: "u3", Total: 25.0},
			},
		},
	}

	for _, tt := range tests {
		res := GenerateInvoices(tt.purchases, tt.taxRate, tt.discountRate)
		if !similar(res, tt.expected) {
			t.Errorf("failed at this case: %v\n", tt.purchases)
		}
	}
}
