package main

import (
	"encoding/json"
	"fmt"
)


// Customer represents a customer structure
// The `omitempty` tag ensures that empty fields are omitted when marshalling to JSON
// SampleOmitEmpty demonstrates the use of `omitempty` in JSON struct tags
func SampleOmitEmpty() {
	// Create a Customer without setting Email
	customer := Customer{ID: 1, Name: "John Doe"}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print JSON
	// 🔹 Email field is omitted because it is empty and has `omitempty`
	fmt.Println(string(jsonData)) // Output: {"id":1,"name":"John Doe"}

	/*
		Tips & Best Practices:
		1. Use `omitempty` to prevent sending unnecessary null or empty values in APIs.
		2. Works for strings, slices, maps, pointers, and numeric types (zero value is omitted).
		3. Helps reduce payload size and makes JSON cleaner for clients.
		4. Combine with `json.MarshalIndent` if you need pretty-printed output.
		5. Fields must be exported (capitalized) to be marshalled.
	*/
}