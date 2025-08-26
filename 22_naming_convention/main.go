package main

import "fmt"

// -------------------------------
// COMMENTS IN GO
// -------------------------------

// Single-line comment example

/*
   This is a block comment.
   Can span multiple lines.
*/

var name = "John Wick" // Inline comment example

// -------------------------------
// NAMING CONVENTIONS IN GO
// -------------------------------

// ✅ Use short, meaningful names in smaller scopes
var _ string   // string
var _ int      // index
var _ int    // number
var _ string // message
var _ string   // value
var _ error  // error value
var _ bool  // status flag (done or not)

// ✅ Prefer camelCase for variables and functions
var _ = 100  // recommended (camelCase)
var _ = 100 // ❌ not recommended (snake_case)

// ✅ Recommended function name
// func writeToFile() {}

// ❌ Not recommended function name
// func write_to_file() {}

// ✅ Acronyms should be uppercase
var _ = true // recommended
var _ = true // ❌ not recommended

func main() {
	// Avoid overly verbose names in small scopes
	var packetsReceived int // ❌ Too verbose
	var n int               // ✅ Short and fine for small scope

	// Suppress unused variable error
	_, _ = packetsReceived, n

	// Print example variable
	fmt.Println("Name:", name)
}
