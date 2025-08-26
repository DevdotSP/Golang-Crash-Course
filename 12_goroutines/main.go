// Using sync.WaitGroup in Go
// --------------------------
// The pattern to use sync.WaitGroup is:
// 1. Create a new variable of type `sync.WaitGroup` (wg).
// 2. Call `wg.Add(n)` where `n` is the number of goroutines you plan to wait for.
// 3. In each goroutine, call `defer wg.Done()` (or wg.Done() before exit)
//    to notify the WaitGroup that the goroutine has completed.
// 4. Call `wg.Wait()` in main() (or any parent function) to block execution
//    until all goroutines have finished.
//
// This allows us to synchronize goroutines and prevent main() from exiting
// before all work is done.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// f1 simulates a goroutine that performs some expensive task.
func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1(goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i =", i)
		// Simulate time-consuming work with sleep.
		time.Sleep(time.Second)
	}

	fmt.Println("f1 execution finished")

	// Step 3:
	// Notify WaitGroup that this goroutine has finished.
	wg.Done()
	// Equivalent: (*wg).Done()
}

// f2 is just a normal function (not a goroutine).
func f2() {
	fmt.Println("f2 execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i =", i)
	}

	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("main execution started")

	// Step 1:
	// Create a WaitGroup to wait for goroutines to complete.
	var wg sync.WaitGroup

	// Step 2:
	// Register one goroutine with wg.Add(1).
	// This tells wg: "I’m going to wait for 1 goroutine to finish".
	wg.Add(1)

	// Launching a goroutine.
	// Pass a pointer to wg so f1 can notify completion.
	go f1(&wg)

	// Print how many goroutines are running right now.
	// Should be 2: (main + f1).
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// Call another function (runs in main, not a goroutine).
	f2()

	// Step 4:
	// Block here until wg counter drops to 0 (meaning all goroutines finished).
	wg.Wait()

	fmt.Println("main execution stopped")
}

// Run the program: go run main.go
//
// EXPECTED OUTPUT (order may slightly vary since goroutines run concurrently):
// main execution started
// No. of Goroutines: 2
// f2 execution started
// f1(goroutine) execution started
// f1, i= 0
// f2(), i= 5
// f2(), i= 6
// f2(), i= 7
// f2 execution finished
// f1, i= 1
// f1, i= 2
// f1 execution finished
// main execution stopped
