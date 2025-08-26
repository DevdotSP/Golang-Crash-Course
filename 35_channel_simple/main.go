package main

import "fmt"

// f1 sends an integer n into the provided channel
func f1(n int, ch chan int) {
	ch <- n // send n into channel (blocks if no receiver ready)
}

func main() {
	// ============================================================
	// 1️⃣ Create a bidirectional channel
	// ============================================================
	// This channel can be used both for sending and receiving.
	c := make(chan int)

	// ============================================================
	// 2️⃣ Create directional channels
	// ============================================================
	// <-chan string : receive-only channel
	// You can only receive from c1, cannot send to it
	c1 := make(<-chan string)

	// chan<- string : send-only channel
	// You can only send to c2, cannot receive from it
	c2 := make(chan<- string)

	// Printing channel types
	fmt.Printf("%T, %T, %T\n", c, c1, c2)
	// Example output: chan int, <-chan string, chan<- string

	// ============================================================
	// 3️⃣ Launch a goroutine that sends a value into the channel
	// ============================================================
	go f1(10, c)

	// ============================================================
	// 4️⃣ Receive the value from the channel
	// ============================================================
	n := <-c
	fmt.Println("Value received:", n)

	fmt.Println("Exiting main...")
}
