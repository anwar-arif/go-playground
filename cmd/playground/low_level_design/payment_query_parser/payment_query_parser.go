package payment_query_parser

import (
	"fmt"
	"log"
	"strings"
)

type PaymentQuery struct {
	User   string
	Amount string
	Status string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var seperators = ":<>"

func extractSeparator(token, key string) string {
	if idx := strings.Index(token, key); idx != -1 {
		sepIdx := idx + len(key)
		if sepIdx < len(token) {
			return string(token[sepIdx])
		}
	}
	return ""
}

func ParsePaymentQuery(query string) (*PaymentQuery, error) {
	tokens := strings.Fields(query)
	if len(tokens) != 3 {
		log.Printf("invalid number of fields %d, expected 3\n", len(tokens))
		return nil, fmt.Errorf("invalid number of fields, expected 3")
	}

	data := make(map[string]string)
	for _, token := range tokens {
		fields := strings.FieldsFunc(token, func(r rune) bool {
			return strings.ContainsRune(seperators, r)
		})
		if len(fields) != 2 {
			log.Printf("skipping %s, key value pair not found\n", token)
			return nil, fmt.Errorf("key value pair not found in %s\n", token)
		}
		key, value := strings.ToLower(fields[0]), fields[1]
		if key == "amount" {
			sep := extractSeparator(token, fields[0])
			value = sep + value
		}
		data[key] = value
	}

	required := []string{"user", "amount", "status"}
	for _, field := range required {
		if _, ok := data[field]; !ok {
			return nil, fmt.Errorf("missing required field: %s\n", field)
		}
	}

	return &PaymentQuery{
		User:   data["user"],
		Amount: data["amount"],
		Status: data["status"],
	}, nil
}

func RunParsePaymentQuery() {
	input := "user:1234 amount100 status:completed"
	res, err := ParsePaymentQuery(input)
	if err != nil {
		log.Printf(err.Error())
	} else {
		fmt.Println(res)
	}
}
