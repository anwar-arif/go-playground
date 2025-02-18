package marko_pro

import "fmt"

// What is the output of this code

func AddN(n int) func(*int) {
	return func(x *int) {
		*x += n
	}
}

func main() {
	a := AddN(1)
	b := AddN(2)
	var c int
	a(&c)
	b(&c)
	fmt.Println(c)
}
