package beginer

import (
	"fmt"

	"golang.org/x/exp/rand"
)

// LuckyNumber represents a single lucky number
type LuckyNumber struct {
	number int
}

// Person contains a slice of LuckyNumber
type Person struct {
	lucky_numbers []LuckyNumber
}

// SampleInitializeStruct demonstrates how to initialize a struct
// and populate a slice of nested structs with random values
func SampleInitializeStruct() {
	// Create a slice of 10 LuckyNumber structs
	tmp := make([]LuckyNumber, 10)

	// Assign a random number (0-99) to each LuckyNumber
	for i := range tmp {
		tmp[i].number = rand.Intn(100)
	}

	// Create a Person instance with the slice of lucky numbers
	a := Person{tmp}

	// Print the Person struct and its lucky numbers
	fmt.Println(a)
}
