package learn

import "fmt"

func ArrayRun() {
	fmt.Println("Running ArrayRun()...")
	sample := [2]int{2, 3}
	fmt.Printf("len %d, array: %v\n", len(sample), sample)
	fmt.Printf("first elem: %d\n", sample[0])
}
