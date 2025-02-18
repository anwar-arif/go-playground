package marko_pro

import "fmt"

// What is the output of this program

func main() {
	i := 0
	ch := make(chan struct{})
	go func() {
		i++
		close(ch)
		i++
	}()
	<-ch
	fmt.Println(i)
}
