package regex_problems

import (
	"fmt"
	"regexp"
	"strings"
)

func MarkSensitiveData(log string) string {
	re := regexp.MustCompile(`\b\d{16}\b`)
	masked := re.ReplaceAllStringFunc(log, func(card string) string {
		return strings.Repeat("*", len(card)-4) + card[len(card)-4:]
	})
	return masked
}

func RunSanitiseLogs() {
	log := "User 1234 made payment with card 4242424242424242"
	fmt.Println(MarkSensitiveData(log))
	fmt.Println(MarkSensitiveData("Card numbers: 1111222233334444 and 5555666677778888"))
}
