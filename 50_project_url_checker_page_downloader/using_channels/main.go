package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// checkAndSaveBody makes an HTTP GET request to the given URL.
// It sends back the result (status message) to the channel `c`.
func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		// If the request fails, send DOWN message into the channel
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s
	} else {
		// If the request succeeds, send UP message with status code
		s := fmt.Sprintf("%s -> Status code: %d \n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		c <- s
	}
}

func main() {
	// List of websites to check
	urls := []string{
		"https://www.golang.org",
		"https://www.google.com",
		"https://www.medium.com",
	}

	// Create a channel to communicate between goroutines
	c := make(chan string)

	// Launch a goroutine for each URL check
	for _, url := range urls {
		go checkAndSaveBody(url, c)
	}

	// Print how many goroutines are currently running
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	/*
		OPTION 1: Finite loop
		You can use this block to wait for responses equal to the number of URLs.
		This prevents infinite looping.
		Example:
			for i := 0; i < len(urls); i++ {
				fmt.Println(<-c)
			}
	*/

	/*
		OPTION 2: Continuous monitoring loop
		This block will keep checking URLs indefinitely at fixed intervals.
		Example:
			for {
				go checkAndSaveBody(<-c, c)
				fmt.Println(strings.Repeat("#", 30))
				time.Sleep(time.Second * 10)
			}
	*/

	// Current implementation: Infinite loop using `range c`
	// Reads messages from the channel `c`
	for url := range c {
		// Launch another goroutine to re-check the URL every 2 seconds
		go func(u string) {
			time.Sleep(time.Second * 2)  // wait before re-checking
			checkAndSaveBody(u, c)       // re-check and send result again
		}(url)
	}
}
