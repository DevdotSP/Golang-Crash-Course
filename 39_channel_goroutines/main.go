package main

import (
	"fmt"
	"strings"
)

// factorial computes the factorial of n and sends the result to a channel
func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	// sending the factorial result into the channel
	// this is a blocking operation until some goroutine receives from the channel
	c <- f
}

func main() {

	// ------------------------------------------------------------
	// 1️⃣ Create a channel for integers
	// - This channel will carry factorial results
	// ------------------------------------------------------------
	ch := make(chan int)

	// defer closing the channel until main() returns
	// ensures resources are cleaned up
	defer close(ch)

	// ------------------------------------------------------------
	// 2️⃣ Launch a single goroutine to compute factorial(5)
	// ------------------------------------------------------------
	go factorial(5, ch)

	// Receiving the result from the channel
	// - This is a blocking operation: main waits here until the goroutine sends a value
	f := <-ch
	fmt.Println("5 factorial =", f)

	// ------------------------------------------------------------
	// 3️⃣ Spawning multiple goroutines (1..20)
	// ------------------------------------------------------------
	// Note: each factorial is sent to the same channel
	// Receiving immediately after starting a goroutine ensures we print the correct value
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Printf("Factorial of %d: %d\n", i, f)
	}

	fmt.Println(strings.Repeat("#", 10))

	// ------------------------------------------------------------
	// 4️⃣ Spawning anonymous goroutines for factorials 5..14
	// ------------------------------------------------------------
	for i := 5; i < 15; i++ {
		// Use a parameter n to avoid closure capturing issues
		go func(n int, c chan int) {
			f := 1
			for j := 2; j <= n; j++ { // note: use j instead of i to avoid shadowing
				f *= j
			}
			// send the factorial result into the channel
			c <- f
		}(i, ch)

		// Receive immediately to avoid race conditions and ensure order
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}
}
