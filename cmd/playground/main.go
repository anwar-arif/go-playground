package main

import (
	"playground/cmd/playground/channels"
	"playground/cmd/playground/learn"
	"playground/cmd/playground/mutex"
	"playground/cmd/playground/wait_group"
)

func main() {
	learn.StructRun()
	learn.ArrayRun()
	channels.RunChannels()
	channels.RunBufferedChannel()
	wait_group.RunWaitGroup()
	mutex.RunMutex()
}
