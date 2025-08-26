package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --- Example 1: Convert string to int using strconv.Atoi ---
	// "45" is a valid numeric string, so it will convert successfully.
	i, err := strconv.Atoi("45")

	// Always check the error returned by Atoi
	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Println("Converted successfully:", i) // prints: 45
	}

	// --- Example 2: Short declaration inside if statement ---
	// Here, we declare i and err inside the if statement.
	// These variables only exist inside this if/else block (scoped locally).
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is", i) // prints: 34
	} else {
		fmt.Println("Conversion failed:", err)
	}

	// --- Example 3: Working with slices of empty interface (interface{}) ---
	// A slice of interface{} can hold values of any type.
	// Start with two strings.
	sample := []interface{}{"apple", "orange"}

	// Append another string to the slice.
	sample = append(sample, "hello")

	// Print the second element in the slice (index 1 = "orange").
	fmt.Println("Second element:", sample[1])

	// Iterate over the slice and print each element.
	// The values can be of any type since interface{} is used.
	for _, value := range sample {
		fmt.Println("Element:", value)
	}
}
