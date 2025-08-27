package main

import (
	"encoding/json"
	"fmt"
)

// SampleJSONarrays demonstrates how to convert a JSON array into Go structs.
func SampleJSONarrays() {
	// Example JSON array string
	// This is commonly what you'd get from a REST API or a database returning JSON
	jsonString := `[{"id":1,"name":"Laptop"},{"id":2,"name":"Phone"}]`

	// Create a slice of Product to hold the unmarshalled data
	var products []Product

	// Unmarshal parses the JSON-encoded data and stores the result in products
	// 🔹 Converts JSON array into a Go slice of structs
	err := json.Unmarshal([]byte(jsonString), &products)
	if err != nil {
		// Handle JSON parsing errors
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Iterate over the slice and access each struct's fields
	for _, product := range products {
		fmt.Println("Product Name:", product.Name)
	}

	/*
		Use-case scenarios:
		1. When fetching JSON data from a database that stores JSON arrays (PostgreSQL JSONB).
		2. When calling an external API that returns JSON arrays (e.g., list of products, users, orders).
		3. When reading JSON files and converting them to Go structs for processing.
		4. Useful in web services to decode incoming JSON request payloads and then respond.
	*/
}
