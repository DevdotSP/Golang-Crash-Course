package main

import (
    "fmt"
    "time"
)

func main() {

    // ============================================================
    // 1️⃣ Record the current timestamp in milliseconds
    // - UnixNano() returns nanoseconds since epoch
    // - Divide by 1_000_000 to convert to milliseconds
    // ============================================================
    start := time.Now().UnixNano() / 1_000_000

    // ============================================================
    // 2️⃣ Select statement overview:
    // - Lets a goroutine wait on multiple channel operations
    // - Blocks until one of its cases is ready
    // - Only works with channels
    // ============================================================

    // ============================================================
    // 3️⃣ Declare two channels of type string
    // ============================================================
    c1 := make(chan string)
    c2 := make(chan string)

    // ============================================================
    // 4️⃣ Start first goroutine that sends a message after 2 seconds
    // ============================================================
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "Hello!" // send message into c1
    }()

    // ============================================================
    // 5️⃣ Start second goroutine that sends a message after 1 second
    // ============================================================
    go func() {
        time.Sleep(1 * time.Second)
        c2 <- "Salut!" // send message into c2
    }()

    // ============================================================
    // 6️⃣ Use select to wait for messages from both channels
    // - Loop runs twice since we expect two messages
    // - Whichever channel receives first will trigger its case
    // ============================================================
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1: // if c1 has a message, receive it
            fmt.Println("Received", msg1)
        case msg2 := <-c2: // if c2 has a message, receive it
            fmt.Println("Received", msg2)
        }
    }

    // ============================================================
    // 7️⃣ Calculate and print total execution time in milliseconds
    // - Demonstrates that goroutines executed concurrently (~2 seconds)
    // ============================================================
    end := time.Now().UnixNano() / 1_000_000
    fmt.Println("Total execution time (ms):", end-start)

    // ============================================================
    // 8️⃣ Notes:
    // - Basic sends/receives are blocking
    // - Using `select` with a `default` clause enables non-blocking channel operations
    // ============================================================
}
