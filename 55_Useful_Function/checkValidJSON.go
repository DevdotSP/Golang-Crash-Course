package main

import (
	"encoding/json" // Provides functions for working with JSON data
	"fmt"           // Provides functions for formatted I/O, like printing to console
)

// CheckValidJSON checks whether a string is valid JSON or not.
func CheckValidJSON() {
	// Example JSON string. This should be in a valid JSON format.
	jsonString := `{"id":1,"name":"John Doe","email":"john@example.com"}`

	// json.Valid checks if the input byte slice is a valid JSON encoding.
	// It returns true if valid, false otherwise.
	if json.Valid([]byte(jsonString)) {
		fmt.Println("Valid JSON") // Print this if JSON is valid
	} else {
		fmt.Println("Invalid JSON") // Print this if JSON is not valid
	}
}

