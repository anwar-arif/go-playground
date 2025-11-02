package regex_problems

import (
	"fmt"
	"regexp"
)

func TemplateVariableExpander(template string, vars map[string]string) string {
	// Match any ${varName} pattern
	re := regexp.MustCompile(`\$\{([a-zA-Z0-9_]+)\}`)

	// Replace each match using a callback
	result := re.ReplaceAllStringFunc(template, func(match string) string {
		// Extract just the variable name using capture group
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // no match — return original
		}
		key := submatches[1]

		// Look up value in map
		if val, ok := vars[key]; ok {
			return val
		}
		// If not found, keep the original placeholder
		return match
	})

	return result
}

func RunTemplateVariableExpander() {
	template := "Payment of ${amount} received for ${user}"
	vars := map[string]string{
		"amount": "120.50",
		"user":   "Alice",
	}

	output := TemplateVariableExpander(template, vars)
	fmt.Println(output)
}
