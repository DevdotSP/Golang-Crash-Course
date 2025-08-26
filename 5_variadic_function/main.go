// =============================================
// Variadic Functions in Go
// =============================================
//
// - A variadic function can accept a variable number of arguments.
// - Defined using an ellipsis (...) before the parameter type.
// - Variadic parameters are internally represented as a slice.
// - A function can only have one variadic parameter, and it must be the last one.
//
// Example: func f(x int, y ...string) {}
// - x is a regular parameter
// - y is variadic, so you can pass 0 or many string values.
//
// =============================================

// 1. Grouped explanation into sections (definition, examples, rules).
// 2. Added why behind each function instead of just what.
// 3. Clarified the behavior of variadic parameters (slice semantics).


package main

import (
    "fmt"
    "strings"
)

// f1: Simple variadic function
// Accepts any number of ints (zero or more).
// Inside the function, "a" is of type []int (a slice of ints).
func f1(a ...int) {
    fmt.Printf("%T\n", a)  // prints the type: []int
    fmt.Printf("%#v\n", a) // prints the actual slice values
}

// f2: Demonstrates that variadic parameters behave like slices.
// Modifying "a[0]" also modifies the underlying slice when shared.
func f2(a ...int) {
    a[0] = 50
}



// sumAndProduct: Variadic function that calculates
// the sum and product of all its float64 arguments.
// Returns (sum, product).
func sumAndProduct(a ...float64) (float64, float64) {
    sum := 0.0
    product := 1.0

    for _, v := range a {
        sum += v
        product *= v
    }

    return sum, product
}

// personInformation: Example of mixing normal and variadic parameters.
// Rules: Non-variadic parameters always come first.
// This function joins all name strings into a full name.
func personInformation(age int, names ...string) string {
    fullName := strings.Join(names, " ")
    return fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)
}

func main() {
    // ================================================
    // Calling variadic functions
    // ================================================

    f1(1, 2, 3, 4) // a becomes []int{1, 2, 3, 4}
    f1()           // a becomes nil slice ([]int(nil))
    f2()

    // Example of a built-in variadic function: append()
    nums := []int{1, 2}
    nums = append(nums, 3, 4) // appends 3 and 4 → []int{1, 2, 3, 4}
    fmt.Println(nums)

    // Call sumAndProduct with 3 numbers
    s, p := sumAndProduct(2., 5., 10.)
    fmt.Println(s, p) // Output: 17 100

    // Call personInformation with multiple names
    info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
    fmt.Println(info) // Output: Age: 35, Full Name: Wolfgang Amadeus Mozart
}
