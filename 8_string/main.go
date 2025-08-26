package main

import (
    "fmt"
)

func main() {

    // 🔹 Strings in Go
    // - Always written inside double quotes "..." or raw backticks `...`
    // - Internally stored as a read-only slice of bytes
    // - Encoded as UTF-8 by default
    // - Immutable (you cannot change a string in place)

    // Declare and initialize a string
    s1 := "Hi there  Go!"

    // Print a string in different formats:
    fmt.Printf("%s\n", s1) // %s → print string as-is -> Hi there  Go!
    fmt.Printf("%q\n", s1) // %q → print with quotes -> "Hi there  Go!"

    // 🔹 Escaping quotes
    // Use backslash `\"` if you need double quotes inside a normal string
    fmt.Println("He say: \"Hello!\"")

    // 🔹 Raw string literal
    // Use backticks `...` when you don’t want escape sequences processed
    fmt.Println(`He say: "Hello!"`)

    // Example: raw string literal (no escaping, everything is taken literally)
    s2 := `Hi there Go!`
    fmt.Println(s2)

    // 🔹 Multiline strings
    // Option 1: Use `\n` inside a normal string to insert a newline
    fmt.Println("Price: 100 \nBrand: Nike")

    // Option 2: Use backticks for natural multiline formatting
    fmt.Println(`
Price: 100
Brand: Nike`)

    // 🔹 File paths
    // With backslashes, raw strings are easier (no double escaping)
    fmt.Println(`C:\Users\Andrei`)   // raw string → backslashes kept as-is
    fmt.Println("C:\\Users\\Andrei") // normal string → needs escaping

    // 🔹 Concatenation
    // Strings are immutable → each + creates a new string in memory
    var s3 = "I love " + "Go " + "Programming"
    fmt.Println(s3 + "!") // -> I love Go Programming!

    // 🔹 Indexing
    // Accessing a character returns the *byte* (uint8), not the character itself
    fmt.Println("Element at index zero:", s3[0]) // -> 73 (ASCII code for 'I')

    // 🔹 Immutability
    // You cannot modify individual characters of a string directly
    // s3[5] = 'x' // ❌ compile-time error
}
