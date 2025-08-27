package beginer

import (
	"log"
	"reflect"
)

// SampleCheckStruct demonstrates how to check if certain fields exist in a struct
func SampleCheckStruct() {
	// Define a sample struct with boolean fields
	type test struct {
		A bool
		B bool
		C bool
	}

	// Create a pointer to a struct instance
	v := new(test)

	// Use reflection to get the value that the pointer points to
	metaValue := reflect.ValueOf(v).Elem()

	// List of field names we want to check
	for _, name := range []string{"A", "C", "Z"} {
		// Get the field by its name
		field := metaValue.FieldByName(name)

		// If the field does not exist, FieldByName returns a zero Value
		if field == (reflect.Value{}) {
			log.Printf("Field %s does not exist in struct", name)
		}
	}
}
