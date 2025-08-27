package beginer

import (
	"encoding/json"
	"fmt"
)

// MyStruct is a sample struct with Name and Score fields
type MyStruct struct {
	Name  string
	Score int
}

// StructToMapString demonstrates how to convert a struct to a map[string]interface{}
// This can be useful for dynamic operations, like JSON processing or flexible data handling
func StructToMapString() {
	// Create an instance of MyStruct
	ms := MyStruct{Name: "John", Score: 34}

	// Declare a map to hold the struct fields as key-value pairs
	var myMap map[string]interface{}

	// Step 1: Marshal the struct to JSON bytes
	data, _ := json.Marshal(ms)

	// Step 2: Unmarshal the JSON bytes into a map
	json.Unmarshal(data, &myMap)

	// Access and print individual fields using map keys
	fmt.Println(myMap["Name"])  // Output: John
	fmt.Println(myMap["Score"]) // Output: 34
}
