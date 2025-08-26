// =============================================
// CONSTANTS IN GO
// =============================================
//
// - A constant is an immutable value known at compile time.
// - Constants can be declared as either:
//     1. Typed:   const a float64 = 5.1
//     2. Untyped: const b = 6.7
//
// - Typed constants have a fixed type.
// - Untyped constants are "kind constants", meaning they can be
//   implicitly converted to compatible types when assigned.
//
// - Constants can participate in compile-time expressions.
// - If the expression result fits the target type, it works.
//
// =============================================

package main

import "fmt"

func main() {
	// -------------------------------
	// Typed Constant
	// -------------------------------
	const a float64 = 5.1 // explicitly typed as float64
	fmt.Println(a)

	// -------------------------------
	// Untyped Constant
	// -------------------------------
	const b = 6.7 // no type → "kind float" (flexible until used)
	fmt.Println(b)

	// Typed constant 'a' can multiply with untyped 'b' → result is float64
	const c float64 = a * b
	fmt.Println(c)

	// String concatenation at compile time
	const str = "Hello " + "Go!"
	fmt.Println(str)

	// Boolean constant from expression (evaluated at compile time)
	const d = 5 > 10
	fmt.Println(d) // false

	// -------------------------------
	// Type Restriction Example
	// -------------------------------
	// ERROR: cannot mix typed int and float64 directly
	// const x int = 5
	// const y float64 = 2.2 * 5

	// Correct approach: use untyped constant
	const x = 5           // untyped → "kind integer"
	const y = 2.2 * 5     // untyped → "kind float"
	fmt.Printf("%T\n", y) // float64

	// -------------------------------
	// Assigning Untyped Constants
	// -------------------------------
	// Untyped constant `x` can adapt to int, float64, or byte
	var i int = x
	var j float64 = x
	var p byte = x       // byte = alias for uint8
	fmt.Println(i, j, p) // Output: 5 5 5

	// -------------------------------
	// Type Inference with Variables
	// -------------------------------
	const r = 5            // untyped constant
	var rr = r             // rr gets type int (default type for integer constants)
	fmt.Printf("%T\n", rr) // int
}
