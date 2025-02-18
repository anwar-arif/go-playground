package main

import "fmt"

// This code snippet is part of a messaging system. The `Process` function requires a `Sender` interface implementation,
// but only a `Send` function is provided, not an interface. How can the `Process` function be called using the `Send` function?

type Message struct{}

type Sender interface {
	Send(Message)
}

func Process(sender Sender) {
	sender.Send(Message{})
}

func Send(message Message) {
	fmt.Println("Message sent")
}

type SendAdapter struct {
}

func (SendAdapter) Send(msg Message) {
	Send(msg)
}
