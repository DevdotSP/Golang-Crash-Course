package main

import (
    "fmt"
)

func main() {

    // ================================
    //   LABEL STATEMENT EXAMPLE
    // ================================

    // Declare a variable named 'outer'
    // Note: Variable names and label names do not conflict.
    outer := 19
    _ = outer // (silence unused variable warning)

    // Define two arrays: a list of people and a list of friends
    people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
    friends := [2]string{"Mark", "Marry"}

    // Goal: Search for the first friend inside the list of people.
    // If found, immediately exit BOTH loops using a label.

outer: // Define a label named "outer"
    for index, name := range people {        // Iterate over people
        for _, friend := range friends {     // Iterate over friends
            if name == friend {
                fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
                break outer // Exit the outer loop (not just the inner loop)
            }
        }
    }

    fmt.Println("Next instruction after the break.")


    // ================================
    //   GOTO STATEMENT EXAMPLE
    // ================================

    // The following code uses `goto` to create a loop,
    // mimicking the behavior of a `for` loop.
    i := 0
loop: // Define a label named "loop"
    if i < 5 {
        fmt.Println(i)
        i++
        goto loop // Jump back to the "loop" label
    }

    // ================================
    //   INVALID GOTO EXAMPLE
    // ================================
    // It’s not allowed to jump over a variable declaration
    // because it would skip initialization.

    // goto todo // ❌ ERROR: cannot jump over the declaration of x
    // x := 5
    // todo:
    // fmt.Println("something here")
}
