package wait_group

import (
	"fmt"
	"sync"
	"time"
)

func RunWaitGroup() {
	fmt.Println("Running RunWaitGroup()...")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// avoid re-use the same reference of i in each goroutine closure
		// https://go.dev/doc/faq#closures_and_goroutines
		workerId := i
		go func(workerId int) {
			defer wg.Done()
			worker(workerId)
		}(workerId)
	}

	wg.Wait()
}

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	// simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}
