package main

import "fmt"

func main() {
	// ===============================
	// BASIC CONSTANTS
	// ===============================

	// Constants must be initialized at declaration
	const days int = 7 // typed constant
	const pi = 3.14    // untyped constant (type inferred later)
    fmt.Println(days,pi)

	// Valid constant kinds: bool, rune, int, float, complex, string

	// Grouped constant declaration
	const (
		a         = 5   // untyped constant
		b float64 = 0.1 // typed constant
	)
    fmt.Println(a,b)

	// Multiple constants in a single line
	const n, m int = 4, 5
    fmt.Println(n,m)

	// If no value is given, Go reuses the last one
	const (
		min1 = -500
		max1 // inherits value = -500
		max2 // inherits again = -500
	)

    fmt.Println(min1,max1,max2)

	// ===============================
	// CONSTANT RULES
	// ===============================

	const temp int = 100
	// temp = 50 // ❌ ERROR: constants cannot be reassigned
    fmt.Println(temp)

	// const power = math.Pow(2, 3) // ❌ ERROR: runtime functions cannot initialize constants

	t := 5
	// const tc = t // ❌ ERROR: variables (runtime) cannot initialize constants

	// ✅ allowed: functions like len() if input is a literal (compile-time known)
	const l1 = len("Hello")
    fmt.Println(l1)

	str := "Hello"
	// const l2 = len(str) // ❌ ERROR: str is a variable (runtime value)

	_, _ = t, str

	// ===============================
	// UNTYPED CONSTANTS
	// ===============================

	const x = 5           // untyped
	const y float64 = 1.1 // typed

	var v1 int = 5
	var v2 float64 = 1.1

	fmt.Println(x * y) // ✅ 5.5 — untyped x adapts to float64 at first use

	// fmt.Println(v1 * v2)
	// ❌ ERROR: mismatched types (int vs float64)

	_, _ = v1, v2

	// ===============================
	// IOTA (CONSTANT COUNTER)
	// ===============================

	// iota is a counter used inside const blocks.
	// It starts at 0 for the first constant in the block,
	// then increments by 1 for each new line.

	const (
		c1 = iota // 0
		c2 = iota // 1
		c3 = iota // 2
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	// More practical: Enums
	const (
		North = iota // 0
		East         // 1
		South        // 2
		West         // 3
	)

    fmt.Println(North,East,South,West)

	// You can also use iota with expressions (patterns)
	const (
		c11 = iota * 2 // 0 * 2 = 0
		c22            // 1 * 2 = 2
		c33            // 2 * 2 = 4
	)

	fmt.Println(c11,c22,c33)
}
