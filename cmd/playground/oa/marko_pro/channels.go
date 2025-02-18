package marko_pro

//Which is not a possible output of this code

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	go func() {
//		v := 1
//		ch1 <- v
//		fmt.Println("A")
//		v2 := <-ch2
//		fmt.Println("B")
//	}()
//
//	for i := 0; i < 2; i++ {
//		select {
//		case ch2 <- i:
//			fmt.Println("C")
//		case <-ch1:
//			fmt.Println("D")
//		}
//	}
//
//	fmt.Println("E", v, v2)
//}
