package diffuc

import (
	"fmt"
	"reflect"
)

// ExampleSlice demonstrates different ways to create, manipulate, slice, and append slices in Go
func ExampleSlice() {
	// Create empty slices and check their type
	CreateEmptySlice()

	// Initialize slices using make() with length and capacity
	InitWithMake()

	// Initialize slices using slice literals
	SliceLiteral()

	// Initialize slices using new() and slicing
	InitWithNew()

	// Change a specific element in a slice
	ChangeItemSlice()

	// Various slice tricks: slicing ranges, last elements, full slices, etc.
	TricksSlice()

	// Append elements to a slice
	AppendSlice()

	// Remove element at a specific index
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println("Original slice:", strSlice)

	strSlice = RemoveIndex(strSlice, 3) // remove element at index 3 ("Germany")
	fmt.Println("After removing index 3:", strSlice)
}

// -------------------------------
// Create empty slices and check type
func CreateEmptySlice() {
	var intSlice []int
	var strSlice []string

	fmt.Println("Type of intSlice:", reflect.ValueOf(intSlice).Kind())
	fmt.Println("Type of strSlice:", reflect.ValueOf(strSlice).Kind())
}

// -------------------------------
// Initialize slices using new()
func InitWithNew() {
	// Create an array of size 50, slice the first 10 elements
	var intSlice = new([50]int)[0:10]

	fmt.Println("Type:", reflect.ValueOf(intSlice).Kind())
	fmt.Printf("Len: %v, Cap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println("Slice contents:", intSlice)
}

// -------------------------------
// Initialize slices using make()
func InitWithMake() {
	var intSlice = make([]int, 10)        // length = 10, capacity = 10
	var strSlice = make([]string, 10, 20) // length = 10, capacity = 20

	fmt.Printf("intSlice -> Len: %v, Cap: %v, Type: %v\n", len(intSlice), cap(intSlice), reflect.ValueOf(intSlice).Kind())
	fmt.Printf("strSlice -> Len: %v, Cap: %v, Type: %v\n", len(strSlice), cap(strSlice), reflect.ValueOf(strSlice).Kind())
}

// -------------------------------
// Initialize slices using literals
func SliceLiteral() {
	var intSlice = []int{10, 20, 30, 40}
	var strSlice = []string{"India", "Canada", "Japan"}

	fmt.Printf("intSlice -> Len: %v, Cap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice -> Len: %v, Cap: %v\n", len(strSlice), cap(strSlice))
}

// -------------------------------
// Modify a specific element in a slice
func ChangeItemSlice() {
	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Println("Original:", strSlice)

	strSlice[2] = "Germany" // change element at index 2
	fmt.Println("Modified:", strSlice)
}

// -------------------------------
// Remove element at a specific index
func RemoveIndex(s []string, index int) []string {
	// Append elements before index with elements after index
	return append(s[:index], s[index+1:]...)
}

// -------------------------------
// Slice tricks: ranges, last elements, full slice, etc.
func TricksSlice() {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	fmt.Printf("Original slice: %v\n", countries)

	fmt.Println("First 2 elements:", countries[:2])
	fmt.Println("Elements 1-2:", countries[1:3])
	fmt.Println("From index 2 to end:", countries[2:])
	fmt.Println("Elements 2-4:", countries[2:5])
	fmt.Println("First 3 elements:", countries[0:3])

	// Access last element
	fmt.Println("Last element:", countries[len(countries)-1])

	// Access last two elements
	fmt.Println("Last two elements:", countries[len(countries)-2:])

	// Print all elements
	fmt.Println("All elements (various ways):", countries[:], countries[0:], countries[0:len(countries)])
}

// -------------------------------
// Append one slice to another
func AppendSlice() {
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	// Append all elements of slice1 to slice2
	slice2 = append(slice2, slice1...)
	fmt.Println("After append:", slice2)
}
