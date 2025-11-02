package receipt_reconciliation

import (
	"github.com/anwar-arif/stripe/transactions"
	"math"
)

func amountKey(amount float64) int64 {
	return int64(math.Round(amount*100) / 100)
}

func populate(receipts []transactions.Transaction, count map[string]map[int64]int, delta int) {
	for _, statement := range receipts {
		if _, ok := count[statement.UserId]; !ok {
			count[statement.UserId] = make(map[int64]int)
		}
		count[statement.UserId][amountKey(statement.Amount)] += delta
	}
}

func FindMismatches(receipts, statements []transactions.Transaction) []string {
	// count[userId][amount] stores the net difference between internal and provider entries
	count := make(map[string]map[int64]int)

	populate(receipts, count, 1)
	populate(statements, count, -1)

	records := make([]string, 0)

	for userId, cnt := range count {
		mismatchFound := false
		for _, c := range cnt {
			if c != 0 {
				mismatchFound = true
				break
			}
		}
		if mismatchFound {
			records = append(records, userId)
		}
	}

	return records
}
