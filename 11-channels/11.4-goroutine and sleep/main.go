package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// A list of URLs to check
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel of strings to communicate results
	c := make(chan string)

	// Launch Goroutines to check the status of each URL concurrently
	for _, link := range links {
		go checkLink(link, c)
	}

	// Continuously read from the channel and re-check URLs after a delay
	for l := range c {
		go func(link string) {
			// Sleep for 5 seconds before re-checking
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// Send an HTTP GET request to the URL
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Send the URL to the channel to indicate it might be down
		return
	}

	fmt.Println(link, "is up!")
	c <- link // Send the URL to the channel to indicate it's up
}