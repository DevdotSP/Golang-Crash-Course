package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

// checkAndSaveBody checks if a given URL is reachable.
// If the response is 200 (OK), it saves the response body to a text file.
// After finishing, it signals the WaitGroup that this goroutine is done.
func checkAndSaveBody(url string, wg *sync.WaitGroup) {

	// Attempt to send GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		// Print the HTTP status code
		fmt.Printf("Status Code: %d  ", resp.StatusCode)

		// Only save response body if the server returned 200 OK
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Generate filename based on domain (after //) + ".txt"
			file := strings.Split(url, "//")[1] + ".txt"

			fmt.Printf("Writing response Body to %s\n", file)

			// Save body to local file
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Mark this goroutine as finished
	wg.Done()
}

func main() {
	// List of URLs to check
	urls := []string{
		"https://www.golang.org",
		"https://www.google1.com", // intentionally invalid
		"https://www.medium.com",
	}

	// 1. Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// 2. Tell the WaitGroup how many goroutines to wait for
	wg.Add(len(urls))

	// 3. Launch one goroutine per URL
	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	// Print current number of goroutines (main + 3 workers)
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// 4. Block main() until all goroutines call wg.Done()
	wg.Wait()
}

// Run: go run main.go
//
// **EXPECTED OUTPUT:**
// No. of Goroutines: 4
// https://www.google1.com is DOWN!
// Status Code: 200  Writing response Body to www.golang.org.txt
// Status Code: 200  Writing response Body to www.medium.com.txt
