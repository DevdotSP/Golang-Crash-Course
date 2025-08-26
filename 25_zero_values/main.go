package main

import "fmt"


// 1. This way, the code is grouped into three clear sections:
// 2. Variable declarations
// 3. Type safety rules
// 4. Zero values


func main() {

	// ================================
	// Variable Declarations in Go
	// ================================

	// Explicit type declaration
	var a int = 10

	// Type inference (Go deduces the type)
	var b = 15.5

	// Short variable declaration (shorthand, type is inferred)
	c := "Gopher"

	// Use Blank Identifier (_) to avoid "declared but not used" errors
	_, _, _ = a, b, c

	// ================================
	// Static and Strong Typing in Go
	// ================================
	// Go is statically and strongly typed:
	// - A variable’s type cannot change after declaration
	// - Different types cannot be assigned to each other directly

	// Examples (will cause errors if uncommented):
	// a = 3.14   // ERROR: cannot assign float64 to int
	// a = b      // ERROR: cannot assign float64 to int

	// ================================
	// Zero Values in Go
	// ================================
	// Uninitialized variables automatically get a "zero value":
	// - int → 0
	// - float64 → 0.0
	// - string → ""
	// - bool → false
	// This ensures every variable always has a valid default value.

	var value int     // zero value: 0
	var price float64 // zero value: 0.0
	var name string   // zero value: ""
	var done bool     // zero value: false

	fmt.Println(value, price, name, done) // Output: 0 0.0 "" false
}
