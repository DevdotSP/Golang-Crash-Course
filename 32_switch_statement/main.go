// =============================================
// SWITCH STATEMENTS IN GO
// =============================================
//
// - A switch is a cleaner alternative to multiple if-else statements.
// - It evaluates an expression once and matches it against case values.
// - Cases don’t fall through by default (unlike C/C++/Java).
// - A `default` case runs if no other case matches.
// - Multiple values can be grouped in a single case.
// - Switch can also be used without an explicit expression,
//   in which case it behaves like a chain of if-else.
//
// =============================================

package main

import "fmt"

func main() {
	// --------------------------------
	// Example 1: Switch with string matching
	// --------------------------------
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python! Indentation instead of curly braces.")
	case "Go", "golang": // multiple values in one case
		fmt.Println("Good for you! You are using curly braces {}.")
	default:
		fmt.Println("Any other language is good to start!")
	}

	// --------------------------------
	// Example 2: Switch without an expression
	// (acts like if-else chain)
	// --------------------------------
	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Unexpected case!")
	}

	// --------------------------------
	// Example 3: Switch as a replacement for long if-else
	// (checking integer values)
	// --------------------------------
	day := 8

	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Error: Invalid day number") // default handles unexpected values
	}
}
