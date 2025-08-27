package beginer

import (
	"errors"
	"fmt"
)

// Global variable to accumulate the multiplication results
var result = 1

// chain multiplies the global result with n
// Demonstrates panic and recover
func chain(n int) {
	// Defer a function to recover from panic
	// This ensures the program does not crash and allows continuation
	defer func() {
		if r := recover(); r != nil {
			// r contains the panic message
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// If n is zero, trigger a panic
	if n == 0 {
		panic(errors.New("cannot multiply a number by zero"))
	} else {
		// Multiply global result by n
		result *= n
		fmt.Println("Output:", result)
	}
}

// ExampleRecoverPanic demonstrates how recover works
// Shows that after a panic, the program continues execution
func ExampleRecoverPanic() {
	fmt.Println("Start multiplication chain:")

	chain(5) // result = 5
	chain(2) // result = 10
	chain(0) // panic triggered, recovered, result unchanged
	chain(8) // result = 80 (continues after recovery)

	fmt.Println("Final result:", result)
}
