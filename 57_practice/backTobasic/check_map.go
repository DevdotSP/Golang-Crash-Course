package backtobasic

import (
	"fmt"
)

// SampleCheckMap demonstrates different ways to work with maps, structs, and formatting
func SampleCheckMap() {
	// Uncomment any function to try it
	// FirstMap()      // Basic map lookup example
	// SecondTry()     // Demonstrates printing multiple types
	ThirdTry()        // Nested struct example
}

// ThirdTry demonstrates working with a slice of structs containing nested structs
func ThirdTry() {
	// Define Example struct with a nested Details struct
	type Example struct {
		Name    string `json:"name"`
		Details struct {
			Name string `json:"name"`
		} `json:"details"`
	}

	// Initialize a slice of Example structs
	examples := []Example{
		{
			Name: "John Doe",
			Details: struct {
				Name string `json:"name"`
			}{
				Name: "John Deep",
			},
		},
		{
			Name: "Jane Doe",
			Details: struct {
				Name string `json:"name"`
			}{
				Name: "Jane Deep",
			},
		},
	}

	// Iterate over the slice and print both the top-level and nested struct values
	for _, iteratedModel := range examples {
		fmt.Printf("Top-level Name: %v | Nested Name: %v\n",
			iteratedModel.Name, iteratedModel.Details.Name)
	}

	// Access the first element in the slice directly
	example := examples[0]
	fmt.Printf("\nDirect access - Name: %v\nNested Name: %v\n",
		example.Name, example.Details.Name)
}

// SecondTry demonstrates formatting and printing multiple types with fmt.Printf
func SecondTry() {
	values := "hello"
	number := 2
	truethness := false
	flooating := 20.00

	// %T prints the type, %v prints the value, %d for int, %s for string, %f for float, %t for bool
	fmt.Printf(`Type of values: %T, Value: %v | Number: %d | String: %s | Float: %f | Bool: %t`+"\n",
		values, values, number, values, flooating, truethness)
}

// FirstMap demonstrates checking existence of keys in a map
func FirstMap() {
	// Initialize a map with boolean values
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	person := "Raphael"

	// Check if the key exists and its value is true
	if val, ok := attended[person]; ok && val {
		fmt.Println(person, "was at the meeting")
	} else {
		fmt.Println(person, "was not at the meeting")
	}
}
