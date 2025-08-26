package main

import (
	"fmt"
	"time"
)

func main() {
	// ============================================================
	// 1️⃣ Create an unbuffered channel
	// ============================================================
	// Unbuffered channels have no internal storage.
	// Sending (c <- value) blocks until another goroutine
	// is ready to receive, and receiving (<-c) blocks until
	// a value is sent.
	c1 := make(chan int) // unbuffered channel

	// ============================================================
	// 2️⃣ Launch a goroutine that sends data to the channel
	// ============================================================
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10 // blocks until the main goroutine receives the data
		fmt.Println("func goroutine after sending data into the channel")
	}(c1) // passing the channel to the anonymous function

	// ============================================================
	// 3️⃣ Main goroutine sleeps before receiving
	// ============================================================
	// Sleeping here demonstrates that the sender blocks until
	// the receiver is ready. If we remove the sleep, the program
	// still works but the order of printed lines may vary.
	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	// ============================================================
	// 4️⃣ Main goroutine receives data from the channel
	// ============================================================
	fmt.Println("main goroutine starts receiving data")
	d := <-c1 // this unblocks the sending goroutine
	fmt.Println("main goroutine received data:", d)

	// ============================================================
	// 5️⃣ Sleep to ensure the goroutine finishes printing
	// ============================================================
	time.Sleep(time.Second)

	/*
		Expected Output:

		main goroutine sleeps for 2 seconds
		func goroutine starts sending data into the channel
		main goroutine starts receiving data
		main goroutine received data: 10
		func goroutine after sending data into the channel
	*/

	// ✅ Key point:
	// Unbuffered channels synchronize sending and receiving.
	// The sender goroutine blocks until the receiver reads the value.
}
