package main

import "fmt"

func main() {

    // ** CONTINUE STATEMENT ** //

    // The `continue` statement skips the rest of the current iteration 
    // and moves control back to the start of the loop.
    // It does NOT stop the loop, just skips to the next iteration.
    //
    // Example: print only even numbers between 1 and 10
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            continue // skip odd numbers, jump to the next iteration
        }
        fmt.Println(i) // only even numbers will be printed
    }

    fmt.Println("----")

    // ** BREAK STATEMENT ** //

    // The `break` statement completely exits (terminates) the current loop.  
    // Unlike `continue` which only skips one iteration, `break` stops the loop entirely.
    //
    // Example: find and print the first 5 numbers divisible by 7
    count := 0
    for i := 1; ; i++ { // infinite loop until break is hit
        if i%7 == 0 {
            fmt.Printf("%d is divisible by 7\n", i)
            count++
        }

        if count == 5 { // stop once 5 numbers are found
            break // breaks out of the loop completely
        }
    }

    // After break, program continues here
    fmt.Println("Loop finished. Program continues...")
}
