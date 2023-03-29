package main

import (
	"playground/cmd/playground/channels"
	"playground/cmd/playground/function"
	"playground/cmd/playground/learn"
	"playground/cmd/playground/mutex"
	_select "playground/cmd/playground/select"
	"playground/cmd/playground/ticker"
	"playground/cmd/playground/wait_group"
)

func main() {
	learn.StructRun()
	learn.ArrayRun()
	channels.RunChannels()
	channels.RunBufferedChannel()
	wait_group.RunWaitGroup()
	mutex.RunMutex()
	ticker.RunTicker()
	_select.RunSelect()
	function.RunVariadicFunc()
}
