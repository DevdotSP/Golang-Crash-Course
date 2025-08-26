// This example demonstrates **buffered channels** in Go.  
// Buffered channels allow sending multiple values into the channel without immediate receivers.
// - The sender blocks only when the buffer is full.
// - The receiver blocks only when the buffer is empty.

package main

import (
    "fmt"
    "time"
)

func main() {
    // ------------------------------------------------------------
    // 1️⃣ Create a buffered channel with capacity 3
    // ------------------------------------------------------------
    c1 := make(chan int, 3)
    fmt.Println("Channel's capacity:", cap(c1)) // => 3

    // ------------------------------------------------------------
    // 2️⃣ Launch a goroutine that sends values into the buffered channel
    // ------------------------------------------------------------
    go func(c chan int) {
        for i := 1; i <= 5; i++ {
            fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
            c <- i // send value into channel; blocks if buffer is full
            fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
        }
        // Close the channel to signal no more values will be sent
        close(c)
    }(c1)

    // ------------------------------------------------------------
    // 3️⃣ Simulate delay before receiving
    // ------------------------------------------------------------
    fmt.Println("main goroutine sleeps 2 seconds")
    time.Sleep(time.Second * 2)

    // ------------------------------------------------------------
    // 4️⃣ Receive values from the buffered channel using range
    // ------------------------------------------------------------
    // range iterates until the channel is closed and all buffered values are received
    for v := range c1 {
        fmt.Println("main goroutine received value from channel:", v)
    }

    // ------------------------------------------------------------
    // 5️⃣ Behavior after channel is closed
    // ------------------------------------------------------------
    // Receiving from a closed channel:
    // - returns the zero-value of the channel's type without blocking
    fmt.Println("Receive from closed channel:", <-c1) // => 0

    // Sending to a closed channel:
    // - will cause a runtime panic
    // c1 <- 10 // uncommenting this line will panic: send on closed channel
}
