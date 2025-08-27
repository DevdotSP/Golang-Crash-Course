package diffuc

import "fmt"

// ExampleChannel demonstrates the use of Go channels and goroutines
func ExampleChannel() {

	// Create a buffered channel of integers with capacity 10
	ch := make(chan int, 10) 

	// Check if the channel is full
	if len(ch) == cap(ch) {
		fmt.Println("Channel is full")
	}

	// Launch a goroutine that sends a value to the channel
	go doSomething(ch) 

	// Receive a value from the channel
	x := <-ch 
	fmt.Println("Received from channel:", x)

	// Demonstrate string channels with multiple goroutines
	stringChannel()
}

// stringChannel demonstrates sending and receiving strings with goroutines
func stringChannel() {
	// Buffered channel of size 1
	ch := make(chan string, 1)

	// Launch multiple goroutines that receive from the channel
	for i := 0; i < 3; i++ {
		go func(id int) {
			msg := <-ch
			fmt.Printf("Goroutine %d received message: %s\n", id, msg)
		}(i)
	}

	// Send a message to the channel
	ch <- "Hello, World!"

	// Receive again to show channel behavior (optional)
	msg := <-ch
	fmt.Println("Main goroutine received:", msg)
}

// doSomething sends an integer value to the provided channel
func doSomething(ch chan int) {
	x := 42
	ch <- x // Send value x into the channel
}
