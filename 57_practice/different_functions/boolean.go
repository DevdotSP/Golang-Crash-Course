package diffuc

import "fmt"

// Booleany demonstrates the use of boolean expressions and conditional statements in Go
func Booleany() {
	age := 40

	// Use if-else if-else to handle multiple conditions
	if age < 30 {
		// Executed if age is less than 30
		fmt.Println("age is less than 30")
	} else if age < 40 {
		// Executed if age is 30 <= age < 40
		fmt.Println("age is less than 40")
	} else {
		// Executed if none of the above conditions are true
		fmt.Println("age is not less than 40")
	}
}
