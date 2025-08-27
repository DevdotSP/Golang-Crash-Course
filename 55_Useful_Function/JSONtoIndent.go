package main

import (
	"encoding/json"
	"fmt"
)

// SampleJSONtoIndent demonstrates how to pretty-print a Go struct as JSON.
func SampleJSONtoIndent() {
	// Example struct
	customer := Customer{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// MarshalIndent converts a Go struct to a JSON byte slice with indentation.
	// 🔹 "" -> prefix for each line
	// 🔹 "  " -> indentation (2 spaces)
	jsonData, err := json.MarshalIndent(customer, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Convert byte slice to string and print
	fmt.Println(string(jsonData))

	/*
		Use-case scenarios:
		1. Pretty-printing JSON for API responses to improve readability.
		2. Logging structs as formatted JSON for debugging.
		3. Generating JSON files or reports that are human-readable.
		4. When returning JSON in tools like Postman or Swagger to make the output easier to inspect.
	*/
}