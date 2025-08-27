package diffuc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// ExampleGoroutines demonstrates creating, synchronizing, and communicating between goroutines
func ExampleGoroutines() {
	// --- Simple Goroutines ---
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	// Sleep to allow the above goroutines to finish
	time.Sleep(10 * time.Second)

	// --- Goroutines with WaitGroup ---
	// Add a count of three, one for each goroutine
	wg.Add(3)
	fmt.Println("Start Goroutines with WaitGroup")

	go AnotherresponseSize("https://www.golangprograms.com")
	go AnotherresponseSize("https://stackoverflow.com")
	go AnotherresponseSize("https://coderwall.com")

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Terminating Program")

	// --- Goroutine with channel and commands ---
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause" // Pause work
	time.Sleep(1 * time.Second)
	command <- "Play"  // Resume work
	time.Sleep(1 * time.Second)
	command <- "Stop"  // Stop goroutine
	wg.Wait()           // Wait for completion

	// --- Fetch values from goroutines using channel ---
	nums := make(chan int) // unbuffered channel
	wg.Add(1)
	go DupliresponseSize("https://www.golangprograms.com", nums)

	// Read the result sent from goroutine via channel
	fmt.Println("Response size:", <-nums)
	wg.Wait()
	close(nums) // Close the channel
}

/*
responseSize: Simple goroutine to fetch a URL and print the response size
*/
func responseSize(url string) {
	fmt.Println("Step1: Fetching", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("Step2: Reading", url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step3: Response length for", url, "is", len(body))
}

// Global WaitGroup for synchronization
var wg sync.WaitGroup

/*
AnotherresponseSize: Similar to responseSize but uses WaitGroup to signal completion
*/
func AnotherresponseSize(url string) {
	defer wg.Done() // Signals WaitGroup that this goroutine is done

	fmt.Println("Step1: Fetching", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println("Step2: Reading", url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step3: Response length for", url, "is", len(body))
}

/*
DupliresponseSize: Fetches URL and sends the response length to an unbuffered channel
*/
func DupliresponseSize(url string, nums chan int) {
	defer wg.Done() // Signal completion

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send the length to the channel
	nums <- len(body)
}

/*
work: Simulates some work in a goroutine
*/
var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println("Working, i =", i)
}

/*
routine: Goroutine that listens to command channel to control execution
- "Pause" pauses the work
- "Play" resumes the work
- "Stop" terminates the goroutine
*/
func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	status := "Play"

	for {
		select {
		case cmd := <-command: // Listen for command from channel
			fmt.Println("Command received:", cmd)
			switch cmd {
			case "Stop":
				return // Exit goroutine
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work() // Execute work if not paused
			}
		}
	}
}
