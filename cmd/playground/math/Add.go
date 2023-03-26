package math

import "fmt"

func init() {
	fmt.Println("in Add init")
}

func Add(a, b int) int {
	return a + b
}
