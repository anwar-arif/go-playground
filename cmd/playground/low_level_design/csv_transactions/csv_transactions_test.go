package csv_transactions

import (
	"fmt"
	"testing"
)

func TestParseTransactions(t *testing.T) {
	tests := []struct {
		lines []string
	}{
		{
			[]string{
				"u1,100,1697820000",
				"u2,not_a_number,1697820050",
				"u3,250,1697821000",
			},
		},
	}

	for _, tt := range tests {
		res := ParseTransactions(tt.lines)
		for _, r := range res {
			fmt.Println(r)
		}
	}
}
