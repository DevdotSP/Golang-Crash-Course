package diffuc

import "fmt"

func SliceArrays() {
	// =========================
	// Arrays
	// =========================
	// Fixed-length array of integers
	var ages = [3]int{25, 25, 25}
	// Fixed-length array of strings
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println("Array of ages:", ages, "Length:", len(ages))
	fmt.Println("Array of names:", names, "Length:", len(names))

	// =========================
	// Slices (dynamic size, built on arrays)
	// =========================
	// Create a slice of integers
	var scores = []int{100, 50, 60}
	// Modify element
	scores[2] = 25
	// Append new element (slice grows automatically)
	scores = append(scores, 85)

	fmt.Println("Scores slice:", scores, "Length:", len(scores))

	// =========================
	// Slice ranges
	// =========================
	// Slices can be obtained from arrays using [start:end]
	rangeOne := names[1:3]  // elements from index 1 up to, but not including 3
	rangeTwo := names[2:]   // elements from index 2 to the end
	rangeThree := names[:3] // elements from start up to index 3

	fmt.Println("Slice range names[1:3]:", rangeOne)
	fmt.Println("Slice range names[2:]:", rangeTwo)
	fmt.Println("Slice range names[:3]:", rangeThree)

	// =========================
	// Copy array to a slice
	// =========================
	// copy requires slices, not arrays, so convert array to slice first
	copied := make([]string, len(names)) // create slice with same length
	copy(copied, names[:])               // convert array to slice using [:]
	fmt.Println("Copied slice from names array:", copied)
}
