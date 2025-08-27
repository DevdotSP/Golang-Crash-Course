package diffuc

import "fmt"

// ExampleGotchas demonstrates common Go "gotchas" using the Gotchas struct
func ExampleGotchas(g Gotchas) {
	fmt.Println("Running ExampleGotchas:")

	// ExampleOne: division with float and error handling
	result, _ := g.ExampleOne()
	fmt.Println("result: ", result)

	// ExampleEntryNilOnMap: accessing a nil map
	g.ExampleEntryNilOnMap()

	// SampleThree: accessing a map with initialized value
	g.SampleThree()

	// SampleLiterals: raw string literals
	g.SampleLiterals()
}

// Gotchas struct
type Gotchas struct{}

// ExampleOne demonstrates a division with mixed types and error handling
func (g *Gotchas) ExampleOne() (float64, error) {
	fmt.Println(g) // Prints the pointer value of Gotchas struct
	var x, y = 13, 3.5

	fmt.Println("ExampleOne output:", float64(x)/y)

	// Handle potential division by zero
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	result := float64(x) / y
	return result, nil
}

// ExampleEntryNilOnMap demonstrates accessing keys in a nil map
func (g *Gotchas) ExampleEntryNilOnMap() {
	var rect map[string]int // nil map (not initialized)

	// Accessing a key in a nil map does NOT panic, it returns the zero value
	fmt.Println("Value for 'height':", rect["height"]) // prints 0
	fmt.Println("Length of map:", len(rect))          // prints 0

	// Checking existence of a key
	value, exists := rect["height"]
	fmt.Println("Value:", value)   // zero value 0
	fmt.Println("Key exists:", exists) // false
}

// SampleThree demonstrates accessing a map with initialized values
func (g *Gotchas) SampleThree() {
	var rect = map[string]int{"height": 10} // initialized map
	fmt.Println(rect["height"])             // prints 10
}

// SampleLiterals demonstrates raw string literals using backticks
func (g *Gotchas) SampleLiterals() {
	// Raw string literal: \t and \n are treated literally
	s := `Go\tJava\nPython`
	fmt.Println(s) // prints: Go\tJava\nPython
}
