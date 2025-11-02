package receipt_reconciliation

import (
	"github.com/anwar-arif/stripe/transactions"
	"testing"
	"time"
)

func Same(got []string, expected []string) bool {
	if len(got) != len(expected) {
		return false
	}
	count := make(map[string]int)
	for _, word := range got {
		count[word]++
	}
	for _, word := range expected {
		count[word]--
	}
	for _, cnt := range count {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func TestFindMismatches(t *testing.T) {
	curTime := time.Now().UnixMilli()
	tests := []struct {
		receipts   []transactions.Transaction
		statements []transactions.Transaction
		expected   []string
	}{
		{
			[]transactions.Transaction{
				{"A", 50.00, curTime},
				{"B", 75.00, curTime},
				{"C", 100.00, curTime},
			},
			[]transactions.Transaction{
				{"A", 50.00, curTime},
				{"B", 80.00, curTime},
				{"D", 40.00, curTime},
			},
			[]string{"B", "C", "D"},
		},
	}

	for _, tt := range tests {
		got := FindMismatches(tt.receipts, tt.statements)
		if !Same(got, tt.expected) {
			t.Errorf("expected %v, got %v\n", tt.expected, got)
		}
	}
}
