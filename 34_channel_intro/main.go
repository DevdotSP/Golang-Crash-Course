package main

import "fmt"

// helper function demonstrating channel communication in a separate goroutine
func sendAndReceive(ch chan int) {
    // send a value into the channel
    ch <- 42
    // receive the value from the channel
    val := <-ch
    fmt.Println("Value received in goroutine:", val)
}

func main() {
    // ------------------------------------------------------------
    // 1️⃣ Declaring a channel variable
    // At this point, it is nil
    // ------------------------------------------------------------
    var ch chan int
    fmt.Println("Nil channel:", ch)

    // ------------------------------------------------------------
    // 2️⃣ Initializing a channel using make()
    // Channels must be initialized before use
    // ------------------------------------------------------------
    ch = make(chan int)
    fmt.Println("Initialized channel:", ch)

    // ------------------------------------------------------------
    // 3️⃣ Creating another channel for sending and receiving
    // ------------------------------------------------------------
    c := make(chan int)

    // ------------------------------------------------------------
    // 4️⃣ Sending a value into the channel
    // - This blocks until another goroutine receives from the channel
    // ------------------------------------------------------------
    // Using goroutine to prevent main from blocking
    go func() {
        c <- 10 // send 10 into channel
    }()

    // ------------------------------------------------------------
    // 5️⃣ Receiving a value from the channel
    // ------------------------------------------------------------
    num := <-c
    fmt.Println("Value received:", num)

    // ------------------------------------------------------------
    // 6️⃣ Close the channel after use
    // - Prevents further sends
    // - Receiving from a closed channel is safe; returns zero value
    // ------------------------------------------------------------
    close(c)

    // ------------------------------------------------------------
    // 7️⃣ Demonstrate safe send and receive using a helper function
    // ------------------------------------------------------------
    ch2 := make(chan int)
    go sendAndReceive(ch2)

    // allow goroutine to complete (simple example)
    val := <-ch2
    fmt.Println("Value received in main:", val)
}
