package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	fmt.Println("Waiting for responses...")

	// true while loop in go
	for l := range ch {
		// function literal / anonymous function
		// l is passed as argument to avoid closure capturing the variable
		// which would lead to all goroutines using the last value of l
		// and hence all checking the same link
		// this is a common gotcha in Go
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, ch) // receive from channel and pass to goroutine
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}
