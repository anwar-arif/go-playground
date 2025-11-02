package csv_transactions

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Transaction struct {
	UserId    string
	Amount    float64
	Timestamp time.Time
}

func GetTransaction(line string) (*Transaction, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return nil, errors.New("invalid transaction")
	}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	userId := parts[0]
	amount, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, errors.New("error parsing the amount")
	}
	timestamp, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, errors.New("error parsing the timestamp")
	}
	t := time.Unix(timestamp, 0)
	return &Transaction{
		UserId:    userId,
		Amount:    amount,
		Timestamp: t,
	}, nil
}

func ParseTransactions(lines []string) []Transaction {
	result := make([]Transaction, 0)
	for _, line := range lines {
		tran, err := GetTransaction(line)
		if err == nil {
			result = append(result, *tran)
		} else {
			log.Printf("skipping %s: %v\n", line, err.Error())
		}
	}
	return result
}
