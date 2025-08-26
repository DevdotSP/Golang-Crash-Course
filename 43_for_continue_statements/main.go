package main

import "fmt"

func main() {
	// Example 1: Skipping odd numbers using continue
	// "continue" means skip the rest of the current loop iteration
	// and move to the next iteration.
	for i := 0; i < 10; i++ {
		if i%2 != 0 { // if i is odd
			continue // skip this iteration
		}
		fmt.Println("Even number:", i)
	}

	fmt.Println("----")

	// Example 2: Skipping multiples of 3
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue // skip printing multiples of 3
		}
		fmt.Println("Not divisible by 3:", i)
	}

	fmt.Println("----")

	// Example 3: Using continue inside nested loops
	// Here, we skip when row == column
	for row := 1; row <= 3; row++ {
		for col := 1; col <= 3; col++ {
			if row == col {
				continue // skip printing diagonal elements
			}
			fmt.Printf("row=%d, col=%d\n", row, col)
		}
	}
}
