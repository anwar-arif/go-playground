package channels

import (
	"fmt"
	"time"
)

// https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html

func RunBufferedChannel() {
	fmt.Println("Running RunBufferedChannel()...")
	// create an unbuffered channel
	baton := make(chan int)

	// First runner to his mark
	go Runner(baton)

	// Start the race
	baton <- 1

	// Give the runners time to race
	time.Sleep(500 * time.Millisecond)
}

func Runner(baton chan int) {
	var newRunner int

	// wait to receive the baton
	runner := <-baton

	// start running around the track
	fmt.Printf("Runner %d running with baton\n", runner)

	// new runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// running around the track
	time.Sleep(100 * time.Millisecond)

	// is the race over
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race over!\n", runner)
		return
	}

	// exchange the baton for the next runner
	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)
	baton <- newRunner
}
