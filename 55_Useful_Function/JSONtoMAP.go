package main

import (
	"encoding/json"
	"fmt"
)

// SampleJSONtoMap demonstrates how to unmarshal JSON into a Go map.
// This is useful when the JSON structure is dynamic or not known in advance.
func SampleJSONtoMap() {
	// Example JSON string (commonly returned from APIs or DB JSON columns)
	jsonString := `{"id":1,"name":"John Doe","email":"john@example.com"}`

	// Use map[string]interface{} to store JSON as key-value pairs
	// 🔹 Keys are strings
	// 🔹 Values are empty interfaces (can hold any type: string, number, bool, array, object)
	var data map[string]interface{}

	// Unmarshal converts JSON string to Go map
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Access data by key
	fmt.Println("Name:", data["name"])   // Output: John Doe
	fmt.Println("Email:", data["email"]) // Output: john@example.com

	// Type assertion is needed if you want to use the value as a concrete type
	if id, ok := data["id"].(float64); ok { // JSON numbers are unmarshalled as float64
		fmt.Println("ID:", int(id)) // Convert to int if needed
	}

	/*
		Use-case scenarios:
		1. When the JSON structure is unknown or dynamic (keys may vary).
		2. Quick prototyping with API responses without defining structs.
		3. Parsing JSON from a database column storing arbitrary JSON data.
		4. When you only need some keys from the JSON and don't want to define full structs.
	*/
}
