package main

import "fmt"

func main() {

    // ------------------------------------------------------------
    // 1️⃣ Declaring and initializing string variables
    // ------------------------------------------------------------
    c := "hello" // declare a string variable c and assign the value "hello"
    v := "world" // declare a string variable v and assign the value "world"

    // ------------------------------------------------------------
    // 2️⃣ Printing values using fmt.Printf
    // ------------------------------------------------------------
    // %v is a general-purpose format specifier that prints the value in a default format
    // fmt.Printf does not add a newline by default
    fmt.Printf("%v %v", c, v) // Output: hello world
}
