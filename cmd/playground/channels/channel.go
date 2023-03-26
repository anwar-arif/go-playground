package channels

import (
	"fmt"
	"net/http"
	"time"
)

func RunChannels() {
	fmt.Println("Running RunChannels()...")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://example.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go CheckLink(link, ch)
	}

	for l := range ch {
		// don't use the same reference variable of one goroutine to another
		// https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7824514#overview
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, ch)
		}(l)
	}
}

func CheckLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}
