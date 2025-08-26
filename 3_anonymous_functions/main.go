package main

import "fmt"

// ------------------------------------------------------------
// Function increment
// ------------------------------------------------------------
// increment takes an integer x as input and returns another function.
// The returned function is a closure: it "remembers" the value of x
// and increments it every time the returned function is called.
func increment(x int) func() int {
    return func() int {
        x++       // increment the captured variable x
        return x  // return the new value
    }
}

func main() {
    // ------------------------------------------------------------
    // 1️⃣ Anonymous function example
    // ------------------------------------------------------------
    // Anonymous functions are functions without a name.
    // Here we declare and call it immediately (IIFE style).
    func(msg string) {
        fmt.Println(msg)
    }("I'm an anonymous function!") // passing a string argument and executing it immediately

    // ------------------------------------------------------------
    // 2️⃣ Using closures
    // ------------------------------------------------------------
    // increment(10) returns a function that increments the internal counter starting from 10
    a := increment(10)

    // Each call to 'a()' increases the internal counter by 1 and returns the new value
    fmt.Println(a()) // 11
    fmt.Println(a()) // 12
    fmt.Println(a()) // 13

    // Note: the state of x is preserved between calls because of the closure
}

// ------------------------------------------------------------
// Instructions to run:
// 1. Save this file as main.go
// 2. Open terminal and navigate to the file location
// 3. Execute: go run main.go
// 4. Observe the output:
//    I'm an anonymous function!
//    11
//    12
//    13
// ------------------------------------------------------------
