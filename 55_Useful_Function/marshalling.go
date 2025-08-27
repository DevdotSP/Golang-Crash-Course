package main

import (
	"encoding/json"
	"fmt"
)



// Customer represents a customer structure
// Struct tags (`json:"fieldname"`) define how fields map to JSON keys
// SampleJSON demonstrates both Marshalling and Unmarshalling in Go
func SampleJSON() {
	// -------------------------
	// MARSHALLING: Go struct -> JSON
	// -------------------------
	customer := Customer{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Marshal struct to JSON bytes
	jsonData, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("Marshalled JSON:", string(jsonData))
	// Output: {"id":1,"name":"John Doe","email":"john@example.com"}

	// -------------------------
	// PRETTY PRINT JSON
	// -------------------------
	jsonPretty, _ := json.MarshalIndent(customer, "", "  ")
	fmt.Println("Pretty JSON:\n", string(jsonPretty))
	/*
		Output:
		{
		  "id": 1,
		  "name": "John Doe",
		  "email": "john@example.com"
		}
	*/

	// -------------------------
	// UNMARSHALLING: JSON -> Go struct
	// -------------------------
	jsonString := `{"id":2,"name":"Alice","email":"alice@example.com"}`
	var newCustomer Customer

	// Unmarshal JSON string into struct
	err = json.Unmarshal([]byte(jsonString), &newCustomer)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Unmarshalled Struct:", newCustomer)
	fmt.Println("Name:", newCustomer.Name) // Access struct fields

	// -------------------------
	// UNMARSHALLING into map
	// -------------------------
	jsonMapString := `{"id":3,"name":"Bob","email":"bob@example.com"}`
	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonMapString), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to map:", err)
		return
	}

	fmt.Println("Unmarshalled Map:", data)
	fmt.Println("Name from map:", data["name"])
	/*
		Notes:
		- JSON numbers are unmarshalled as float64 when using map[string]interface{}
		- Requires type assertion to convert to int or other types
	*/
}