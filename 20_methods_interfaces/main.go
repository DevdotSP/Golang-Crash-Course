package main

import (
    "fmt"
    "time"
)

// declaring a new custom type "names"
// it is simply a slice of strings, but we give it a new name
// so we can define methods (functions) on it
type names []string

// defining a method for the "names" type
// (n names) is the receiver → the variable on which this method is called
func (n names) print() {
    // The receiver 'n' is like 'this' or 'self' in OOP languages.
    // Here, 'n' is a copy of the 'names' slice on which we call .print()

    // Iterate over the slice and print index and value
    for i, name := range n {
        fmt.Println(i, name)
    }
}

func main() {

    // Go does not have classes, but it allows methods to be attached to types.
    // A "method set" is the collection of methods defined on a type.

    // Example: Working with built-in type time.Duration
    const day = 24 * time.Hour

    // %T prints the type of the variable
    fmt.Printf("%T\n", day) // Output: time.Duration

    // time.Duration has methods defined on it, e.g. Seconds()
    seconds := day.Seconds()

    // Seconds() converts the duration into a floating-point number (seconds)
    fmt.Printf("%T\n", seconds)               // Output: float64
    fmt.Println("Seconds in a day:", seconds) // Output: 86400

    // Example: Working with our custom "names" type
    friends := names{"Dan", "Marry"}

    // Call the print() method → this prints each element of the slice
    friends.print()

    // Alternative way to call a method:
    // We can explicitly pass the receiver, but it's uncommon in practice
    names.print(friends)
}
