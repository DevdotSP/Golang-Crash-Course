// ============================================================
// ** IMPORTANT **
// ============================================================
// 1️⃣ Run this program on your local machine (not in Go Playground)
//    Execute: go run main.go
//
// 2️⃣ Use Go Race Detector to check for data races
//    Execute: go run -race main.go
//
// This program demonstrates safe concurrent access to a shared variable
// using sync.Mutex to prevent data races.
//
// ============================================================

package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    const gr = 100 // number of goroutines for incrementing and decrementing
    var wg sync.WaitGroup
    wg.Add(gr * 2) // each loop spawns 2 goroutines, so total = gr*2

    // shared variable
    var n int = 0

    // 1️⃣ Mutex declaration
    // A Mutex (Mutual Exclusion) ensures that only one goroutine
    // can access the critical section at a time
    var m sync.Mutex

    for i := 0; i < gr; i++ {

        // -------------------------
        // Goroutine 1: Increment n
        // -------------------------
        go func() {
            // simulate work
            time.Sleep(time.Second / 10)

            // 2️⃣ Lock the mutex before accessing shared variable
            m.Lock()
            n++ // critical section: increment

            // 3️⃣ Unlock after done
            m.Unlock()

            // mark this goroutine as done
            wg.Done()
        }()

        // -------------------------
        // Goroutine 2: Decrement n
        // -------------------------
        go func() {
            time.Sleep(time.Second / 10)

            // Using defer to unlock after decrement
            m.Lock()
            defer m.Unlock()
            n-- // critical section: decrement

            wg.Done()
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()

    // print final value
    fmt.Println(n) // ✅ Expected: 0, increments and decrements cancel out
}
