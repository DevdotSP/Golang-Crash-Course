package diffuc

import (
	"fmt"
	"time"
)

// SwitchCase demonstrates different ways to use switch statements in Go
func SwitchCase() {
	Schedule()              // Simple switch on a value
	DatDay()                // Using fallthrough to execute multiple cases
	ConditionalStatementsu() // Switch with conditional expressions
	MultipleValuesCase()    // Using multiple values in a single case
	TypeSwitchExample()     // Using type switch to check variable type
}

// ----------------------
// Simple switch statement on a value
// ----------------------
func Schedule() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default: // Default executes if no case matches
		fmt.Println("No information available for that day.")
	}
}

// ----------------------
// Switch with fallthrough
// fallthrough executes the next case even if its condition is false
// ----------------------
func DatDay() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// ----------------------
// Switch with conditional expressions (like if-else)
// ----------------------
func ConditionalStatementsu() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// ----------------------
// Switch with multiple values in a single case
// ----------------------
func MultipleValuesCase() {
	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend! Relax!")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Weekday: Time to work.")
	default:
		fmt.Println("Unknown day")
	}
}

// ----------------------
// Type switch: Determine the type of an interface variable
// ----------------------
func TypeSwitchExample() {
	var x interface{} = 25 // Can be any type

	switch v := x.(type) {
	case int:
		fmt.Println("x is an integer:", v)
	case float64:
		fmt.Println("x is a float:", v)
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
