package diffuc

import (
	"fmt"
	"reflect"
)

// VariadicFunction demonstrates usage of variadic functions in Go
func VariadicFunction() {
	// ----------------------
	// Example 1: Variadic function with strings
	// ----------------------
	variadicExample("red", "blue", "green", "yellow")

	// ----------------------
	// Example 2: Variadic function for calculations
	// ----------------------
	// Rectangle: Multiply all values to get area
	fmt.Println("Rectangle Area:", calculation("Rectangle", 20, 30))
	// Square: Only one value considered, area = value * value
	fmt.Println("Square Area:", calculation("Square", 20))

	// ----------------------
	// Example 3: Variadic function with different types (interface{})
	// ----------------------
	variadicExample2(
		1,
		"red",
		true,
		10.5,
		[]string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13},
	)

	// ----------------------
	// Example 4: Using defer to execute a function after current function finishes
	// ----------------------
	defer second() // executes at the end of VariadicFunction
	first()        // executes immediately
}

// ----------------------
// Example 1: Simple variadic function with strings
// ----------------------
func variadicExample(s ...string) {
	fmt.Println("First element:", s[0])
	fmt.Println("Fourth element:", s[3])
}

// ----------------------
// Example 2: Variadic function for calculations
// Accepts string type and variadic integers
// ----------------------
func calculation(shape string, values ...int) int {
	area := 1

	for _, val := range values {
		if shape == "Rectangle" {
			area *= val // multiply all dimensions
		} else if shape == "Square" {
			area = val * val // square only uses one value
		}
	}
	return area
}

// ----------------------
// Example 3: Variadic function accepting any type using interface{}
// ----------------------
func variadicExample2(items ...interface{}) {
	for _, v := range items {
		fmt.Printf("%v -- Type: %v\n", v, reflect.TypeOf(v))
	}
}

// ----------------------
// Example 4: Demonstrating defer
// ----------------------
func first() {
	fmt.Println("First executes immediately")
}

func second() {
	fmt.Println("Second executes last (deferred)")
}
