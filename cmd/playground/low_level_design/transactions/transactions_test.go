package transactions

import (
	"testing"
	"time"
)

func SameUsers(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	count := make(map[string]int, 0)

	for _, name := range first {
		count[name]++
	}
	for _, name := range second {
		count[name]--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestDetectFraudulentUsers(t *testing.T) {
	t0 := time.Second
	tests := []struct {
		transactions []Transaction
		Threshold    float64
		Window       time.Duration
		expected     []string
	}{
		{
			transactions: []Transaction{
				{UserId: "u1", Amount: 400, TimeStamp: t0.Milliseconds()},
				{UserId: "u1", Amount: 350, TimeStamp: t0.Milliseconds() + (2 * time.Minute.Milliseconds())},
				{UserId: "u1", Amount: 300, TimeStamp: t0.Milliseconds() + (4 * time.Minute.Milliseconds())},
			},
			Threshold: 1000,
			Window:    5 * time.Second,
			expected:  []string{"u1"},
		},
	}

	for _, tt := range tests {
		res := DetectFraudulentUsers(tt.transactions, tt.Threshold, tt.Window)
		if !SameUsers(res, tt.expected) {
			t.Errorf("expected %v, got %v\n", tt.expected, res)
		}
	}
}
