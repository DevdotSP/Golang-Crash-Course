package diffuc

import (
	"fmt"
	"reflect"
)

// ArrayExample demonstrates multiple operations with arrays in Go
func ArrayExample() {
	CheckType()          // Check array type using reflection
	ToAccessArray()      // Assign and access array elements
	SelectSpecificElem() // Assign values to specific indices
	FilteringArrays()    // Slice and filter arrays
}

// -------------------- CHECK TYPE --------------------
// Demonstrates how to check if a variable is an array and use a helper function
func CheckType() {
	var intArray [5]int          // Integer array of size 5
	var strArray [5]string       // String array of size 5
	strArray1 := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}

	// Check if an item exists in an array
	fmt.Println(ItemExists(strArray1, "Canada")) // true
	fmt.Println(ItemExists(strArray1, "Africa")) // false

	// Use reflection to get the kind of variable
	fmt.Println(reflect.ValueOf(intArray).Kind()) // Output: array
	fmt.Println(reflect.ValueOf(strArray).Kind()) // Output: array
}

// -------------------- ACCESSING ARRAY --------------------
// Demonstrates assigning and accessing array elements
func ToAccessArray() {
	var theArray [4]string
	theArray[0] = "India"  // Assign value to first element
	theArray[1] = "Canada" // Assign value to second element
	theArray[2] = "Japan"  // Assign value to third element

	leng := len(theArray) // Get length of array
	fmt.Println(leng)

	// Access individual elements
	fmt.Println(theArray[0])
	fmt.Println(theArray[1])
	fmt.Println(theArray[2])
}

// -------------------- SELECT SPECIFIC ELEMENT --------------------
// Demonstrates assigning values to specific indices
func SelectSpecificElem() {
	x := [5]int{1: 10, 3: 30} // Only indices 1 and 3 are initialized
	fmt.Println(x)             // Output: [0 10 0 30 0]
}

// -------------------- ITEM EXISTS --------------------
// Generic helper to check if a specific item exists in an array
func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type: not an array")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// -------------------- FILTERING ARRAYS / SLICING --------------------
// Demonstrates slicing arrays and selecting specific ranges
func FilteringArrays() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("All Countries: %v\n", countries)

	// Slice from start to index 2 (exclusive)
	fmt.Printf("[:2] -> %v\n", countries[:2])

	// Slice from index 1 to 3 (exclusive)
	fmt.Printf("[1:3] -> %v\n", countries[1:3])

	// Slice from index 2 to the end
	fmt.Printf("[2:] -> %v\n", countries[2:])

	// Slice from index 2 to 5 (exclusive)
	fmt.Printf("[2:5] -> %v\n", countries[2:5])

	// Slice from start to index 3 (exclusive)
	fmt.Printf("[0:3] -> %v\n", countries[0:3])

	// Access the last element
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	// Full array slices
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	// Last two elements
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
}
