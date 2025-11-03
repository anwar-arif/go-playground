package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'calculate_fees' function below.
 *
 * The function accepts STRING transaction_csv as parameter.
 */

type TransactionType string
type PaymentProvider string
type PaymentStatus string

var (
	Completed       PaymentStatus = "completed"
	Failed          PaymentStatus = "failed"
	Pending         PaymentStatus = "pending"
	DisputeWon      PaymentStatus = "dispute_won"
	DisputeLost     PaymentStatus = "dispute_lost"
	RefundCompleted PaymentStatus = "refund_completed"
	RefundFailed    PaymentStatus = "refund_failed"
	RefundPending   PaymentStatus = "refund_pending"
)

var (
	Klarna PaymentProvider = "klarna"
	Card   PaymentProvider = "card"
)

var (
	Payment TransactionType = "payment"
	Refund  TransactionType = "refund"
	Dispute TransactionType = "dispute"
)

type Transaction struct {
	Id              string
	Reference       string
	Amount          int
	Currency        string
	Date            string
	MerchantId      string
	BuyerCountry    string
	TransactionType string
	PaymentProvider string
	Status          string
}

type Fee struct {
	Id              string
	TransactionType string
	PaymentProvider string
	Fee             int
}

func (f *Fee) Stringify() string {
	return fmt.Sprintf("%v,%v,%v,%v", f.Id, f.TransactionType, f.PaymentProvider, f.Fee)
}

func CalculatePaymentFee(transaction Transaction) int {

	if transaction.Status == "payment_completed" {
		return int(float64(transaction.Amount)*.021) + 30
	} else if transaction.Status == "dispute_lost" {
		return 15
	} else if transaction.Status == "dispute_won" {
		if transaction.PaymentProvider == "card" {
			return 15
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func ParseTransaction(line string) Transaction {
	parts := strings.Split(line, ",")

	fields := make(map[string]string)
	for index, part := range parts {
		fields[header[index]] = part
	}

	var transaction Transaction
	for key, value := range fields {
		if key == "id" {
			transaction.Id = value
		} else if key == "reference" {
			transaction.Reference = value
		} else if key == "amount" {
			amount, _ := strconv.Atoi(value)
			transaction.Amount = amount
		} else if key == "currency" {
			transaction.Currency = value
		} else if key == "date" {
			transaction.Date = value
		} else if key == "merchant_id" {
			transaction.MerchantId = value
		} else if key == "buyer_country" {
			transaction.BuyerCountry = value
		} else if key == "payment_provider" {
			transaction.PaymentProvider = value
		} else if key == "transaction_type" {
			transaction.TransactionType = value
		} else if key == "status" {
			transaction.Status = value
		}
	}

	return transaction
}

var header []string

func parseHeaders(line string) {
	parts := strings.Split(line, ",")

	for _, fieldName := range parts {
		header = append(header, fieldName)
	}
}

func initialise() {
	header = make([]string, 0)
}

func calculate_fees(transaction_csv string) {
	initialise()
	// Write your code here
	lines := strings.Split(transaction_csv, "\n")
	for index, line := range lines {
		if index == 0 {
			parseHeaders(line)
			continue
		}
		transaction := ParseTransaction(line)
		fee := Fee{
			Id:              transaction.Id,
			TransactionType: transaction.TransactionType,
			PaymentProvider: transaction.PaymentProvider,
			Fee:             CalculatePaymentFee(transaction),
		}
		log.Printf("fee: %v\n", fee.Stringify())
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	file, err := os.Open("./input.csv")
	if err != nil {
		log.Printf("cannot read file: %v\n", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	content := ""
	for scanner.Scan() {
		if len(content) > 0 {
			content += "\n"
		}
		content += scanner.Text()
	}

	log.Printf("content: %v", content)

	calculate_fees(content)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
