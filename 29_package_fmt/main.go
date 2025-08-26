package main

// Package fmt implements formatted I/O functions, similar to C's printf and scanf.
// Commonly used to print output to the console.
import "fmt"

func main() {

    // ---------------------------
    // fmt.Println()
    // ---------------------------
    // Prints arguments separated by spaces.
    // Always adds a space between values and appends a newline at the end.
    fmt.Println("Hello Go World!") // Output: Hello Go World!

    var name, age = "Andrei", 35
    fmt.Println(name, "is", age, "years old.") // Output: Andrei is 35 years old.

    // ---------------------------
    // fmt.Printf()
    // ---------------------------
    // Prints according to a format specifier (verb).
    // Unlike Println, it does NOT automatically add a newline.
    //
    // Common verbs:
    // %d → decimal (int)
    // %f → float (default 6 decimals)
    // %.nf → float with n decimal places
    // %s → string
    // %q → double-quoted string
    // %v → value in default format
    // %#v → Go-syntax representation (useful for debugging)
    // %T → type of the value
    // %t → boolean
    // %p → pointer (address in memory, base 16)
    // %c → character (Unicode code point)
    // %b → binary
    // %x → hexadecimal

    a, b, c := 10, 15.5, "Gophers"
    grades := []int{10, 20, 30}

    // Basic formatting
    fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)
    // Output: a is 10, b is 15.500000, c is Gophers

    // Print with quotes
    fmt.Printf("%q\n", c) // Output: "Gophers"

    // Print slice in default format
    fmt.Printf("%v\n", grades) // Output: [10 20 30]

    // Print slice with Go syntax representation
    fmt.Printf("%#v\n", grades) // Output: []int{10, 20, 30}

    // Print type information
    fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
    // Output: b is of type float64 and grades is of type []int

    // Print memory address of variable
    fmt.Printf("The address of a: %p\n", &a)
    // Example Output: The address of a: 0xc000016128

    // Print Unicode characters from integer values
    fmt.Printf("%c and %c\n", 100, 51011)
    // Output: d and 읃  (characters for Unicode code points 100 and 51011)

    // Format float with precision
    const pi float64 = 3.14159265359
    fmt.Printf("pi is %.4f\n", pi) // Output: pi is 3.1416

    // Print numbers in other bases
    fmt.Printf("255 in base 2 is %b\n", 255)  // Output: 255 in base 2 is 11111111
    fmt.Printf("101 in base 16 is %x\n", 101) // Output: 101 in base 16 is 65

    // ---------------------------
    // fmt.Sprintf()
    // ---------------------------
    // Works like Printf but returns a formatted string instead of printing it.
    s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)

    // Print the formatted string
    fmt.Println(s)
    // Output: a is 10, b is 15.500000, c is Gophers
}
