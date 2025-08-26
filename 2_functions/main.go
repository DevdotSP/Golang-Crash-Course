package main

import (
	"fmt"
	"math"
)

//
// 🔹 Function Definitions
//

// f1 demonstrates a function with **no parameters** and **no return value**.
// It simply prints a message.
func f1() {
	fmt.Println("This is f1() function")
}

// f2 demonstrates a function with **two parameters (int)** and **no return value**.
// It prints the sum of a and b.
// Note: parameters are local to the function.
func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// f3 demonstrates the use of **shorthand notation** for parameters.
// Multiple parameters of the same type can be grouped together.
// This function accepts integers, floats, and a string.
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// f4 demonstrates a function that **takes a single float64 parameter**
// and **returns a single float64 result**.
// It uses math.Pow to calculate "a raised to the power of a".
func f4(a float64) float64 {
	return math.Pow(a, a) // statements after return will never run
}

// f5 demonstrates a function that **returns multiple values**.
// Here, it returns the product and sum of a and b.
func f5(a, b int) (int, int) {
	return a * b, a + b
}

// sum demonstrates the use of a **named return value** (s).
// - s is declared in the function signature and starts with the zero value (0).
// - assigning s = a + b and using a "naked return" implicitly returns s.
func sum(a, b int) (s int) {
	fmt.Println("Initial s:", s) // -> s is initially 0
	s = a + b
	return // naked return, returns the named variable "s"
}

//
// 🔹 Main Function (Program Entry Point)
//
func main() {
	// Example 1: Call a simple function with no params
	f1() // Output: This is f1() function

	// Example 2: Call a function with two parameters
	f2(2, 3) // Output: Sum: 5

	// Example 3: Shorthand parameters of different types
	f3(3, 4, 5, 4.0, 5.5, "hello") // Output: 3 4 5 4 5.5 hello

	// Example 4: Function with return value
	fmt.Println(f4(5.1)) // Output: 5.1^5.1 (approx 3455.94)

	// Example 5: Function with multiple return values
	product, total := f5(6, 7)
	fmt.Printf("Product: %d, Total: %d\n", product, total) // Output: Product: 42, Total: 13

	// Example 6: Function with named return value
	ss := sum(4, 5)
	fmt.Println("Sum result:", ss) // Output: 9

	// Note: Go does not support default parameter values or function overloading.
}
