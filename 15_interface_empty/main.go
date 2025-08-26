// Empty Interface and Interfaces in Go
// ------------------------------------
// In Go, an empty interface (interface{}) is a special type that can hold any value.
// This is because every type in Go implicitly implements the empty interface.
// Interfaces in general allow us to define behavior (methods) that types can implement.
// This code demonstrates:
//   1. How an empty interface works.
//   2. How to use type assertions with empty interfaces.
//   3. How to use custom interfaces to enforce behavior (e.g., Sound with Bark()).

package main

import "fmt"

// Declaring an empty interface (interface with no methods).
// Any type in Go satisfies this interface.
type emptyInterface interface{}

// Declaring a struct with a field of type empty interface.
// This means the field can store any type of value.
type person struct {
	info interface{}
}

// Sound interface declares a behavior (contract) with one method: Bark() string.
type Sound interface {
	Bark() string
}

// Dog struct with a name field.
type Dog struct {
	name string
}

// Cat struct with a name field.
type Cat struct {
	name string
}

// Dog implements Sound interface by defining the Bark() method.
func (d Dog) Bark() string {
	return d.name + " says Woof!"
}

// Cat also implements Sound interface by defining the Bark() method.
func (c Cat) Bark() string {
	return c.name + " says Meow!"
}

// makeSound accepts any type that implements the Sound interface
// and calls its Bark() method.
func makeSound(s Sound) {
	fmt.Println(s.Bark())
}

func main() {
	// Example 1: Using the empty interface
	var empty interface{} // Declares a variable of type empty interface.

	// Assigning different types to the empty interface.
	empty = 5
	fmt.Println(empty) // => 5

	empty = "Go"
	fmt.Println(empty) // => Go

	empty = []int{2, 34, 4}
	fmt.Println(empty) // => [2 34 4]

	// Attempting to use len() directly on empty would fail,
	// because the compiler doesn’t know the underlying type.
	// fmt.Println(len(empty)) -> ERROR

	// To access the actual value, we need a type assertion.
	fmt.Println(len(empty.([]int))) // => 3

	// Example 2: Using empty interface in a struct
	you := person{}

	// Assigning different values to the info field (any type is allowed).
	you.info = "Your name"
	you.info = 40
	you.info = []float64{4.5, 6.0, 8.1}

	// Prints the last assigned value
	fmt.Println(you.info) // => [4.5 6 8.1]

	// Example 3: Interfaces with behavior
	cat := Cat{name: "Muning"}
	dog := Dog{name: "Blackie"}

	// Both Cat and Dog implement the Sound interface,
	// so we can pass them to makeSound().
	makeSound(cat) // => Muning says Meow!
	makeSound(dog) // => Blackie says Woof!
}
