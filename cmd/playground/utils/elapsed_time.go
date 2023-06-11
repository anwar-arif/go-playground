package utils

import (
	"fmt"
	"time"
)

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed time: ", elapsed)
	return elapsed
}

func RunElapsedTime() {
	defer TrackTime(time.Now())
	time.Sleep(500 * time.Millisecond)
}
