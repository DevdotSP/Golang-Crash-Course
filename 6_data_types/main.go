package main

import "fmt"

func main() {

	// ======================
	// 1️⃣ NUMERIC TYPES
	// ======================

	// Unsigned integers (no negative values)
	// uint8   : 0 to 255
	// uint16  : 0 to 65535
	// uint32  : 0 to 4294967295
	// uint64  : 0 to 18446744073709551615

	// Signed integers
	// int8    : -128 to 127
	// int16   : -32768 to 32767
	// int32   : -2147483648 to 2147483647
	// int64   : -9223372036854775808 to 9223372036854775807

	// platform-dependent integers
	// uint    : same size as int, either 32 or 64 bits
	// int     : same size as uint

	// Floating-point numbers
	// float32    : 32-bit IEEE-754 floating point
	// float64    : 64-bit IEEE-754 floating point

	// Complex numbers
	// complex64  : float32 real + imag
	// complex128 : float64 real + imag

	// Aliases
	// byte  : alias for uint8
	// rune  : alias for int32 (represents a Unicode code point)

	// ✅ Examples:
	var i1 int8 = -128     // minimum value for int8
	fmt.Printf("%T\n", i1) // int8

	var i2 uint16 = 65535  // maximum value for uint16
	fmt.Printf("%T\n", i2) // uint16

	var i3 int64 = -324_567_345  // underscores improve readability
	fmt.Printf("%T\n", i3)       // int64
	fmt.Printf("i3 is %d\n", i3) // underscores are ignored

	var f1, f2, f3 float64 = 1.1, -.2, 5. // floats
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var r rune = 'f'      // rune is int32
	fmt.Printf("%T\n", r) // int32
	fmt.Printf("%x\n", r) // hexadecimal Unicode code for 'f'
	fmt.Printf("%c\n", r) // prints the character itself

	var b bool = true
	fmt.Printf("%T\n", b) // bool

	var s string = "Hello Go!"
	fmt.Printf("%T\n", s) // string

	// ======================
	// 2️⃣ COMPOSITE TYPES
	// ======================

	// Array: fixed-size sequence of elements
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers) // [4]int

	// Slice: dynamically-sized sequence of elements
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // []string

	// Map: key-value pairs
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances) // map[string]float64

	// Struct: collection of fields
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)      // main.Person
	fmt.Printf("%T\n", you.name) // main.Person
	fmt.Printf("%T\n", you.age)  // main.Person

	// Pointer: holds memory address of a value
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // *int

	// Function type
	fmt.Printf("%T\n", f) // func()
}

// simple function
func f() {
}
