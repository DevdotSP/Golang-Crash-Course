package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ListOfMethods demonstrates common Go operations, including:
// - String manipulation
// - Slices & arrays
// - Maps
// - File handling
// - Error handling
// - Concurrency
// - Type conversion
// - Time/date handling
// - Structs & methods
func ListOfMethods() {

	// 1. STRING MANIPULATION
	// 🔹 Concatenation using +
	str := "Hello" + " " + "World"
	fmt.Println(str) // Output: Hello World

	// 🔹 Get string length
	fmt.Println(len("Hello")) // Output: 5

	// 🔹 Split string into slice
	strList := "apple,banana,orange"
	parts := strings.Split(strList, ",")
	fmt.Println(parts) // Output: [apple banana orange]

	// 🔹 Join slice of strings into a single string
	words := []string{"Go", "is", "awesome"}
	sentence := strings.Join(words, " ")
	fmt.Println(sentence) // Output: Go is awesome

	// 🔹 Check if substring exists
	fmt.Println(strings.Contains("golang", "go")) // Output: true

	// 🔹 Replace all occurrences of a substring
	newStr := strings.ReplaceAll("hello world", "world", "Go")
	fmt.Println(newStr) // Output: hello Go

	// 🔹 Convert case
	fmt.Println(strings.ToUpper("golang")) // Output: GOLANG
	fmt.Println(strings.ToLower("GOLANG")) // Output: golang

	// 2. SLICES & ARRAYS
	// 🔹 Initialize a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 🔹 Append elements to a slice
	numbers1 := []int{1, 2, 3}
	numbers1 = append(numbers1, 4, 5)
	fmt.Println(numbers1) // Output: [1 2 3 4 5]

	// 🔹 Copy slices
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(destination) // Output: [1 2 3]

	// 3. MAPS (Key-Value Pairs)
	// 🔹 Create a map
	person := map[string]string{
		"name": "Alice",
		"age":  "25",
	}
	fmt.Println(person)

	// 🔹 Add or access values
	person["city"] = "New York"
	fmt.Println(person["name"]) // Output: Alice

	// 🔹 Check if a key exists
	value, exists := person["age"]
	if exists {
		fmt.Println("Age:", value)
	} else {
		fmt.Println("Age not found")
	}

	// 4. FILE HANDLING
	// 🔹 Write to a file
	err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
	if err != nil {
		panic(err)
	}

	// 🔹 Read from a file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// 5. ERROR HANDLING
	result, err := divide(10, 0) // division by zero example
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 6. CONCURRENCY
	// 🔹 Using Goroutines
	go tulogmoTO()
	time.Sleep(1 * time.Second) // wait for goroutine to finish

	// 🔹 Using Channels to communicate between goroutines
	ch := make(chan string)
	go sendMessage(ch)
	message := <-ch
	fmt.Println(message)

	// 7. TYPE CONVERSION
	// 🔹 String to Integer
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(num) // Output: 123
	}

	// 🔹 Integer to String
	str = strconv.Itoa(123)
	fmt.Println(str) // Output: "123"

	// 🔹 String to Float
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f) // Output: 3.14

	// 8. TIME AND DATE HANDLING
	// 🔹 Get current time
	now := time.Now()
	fmt.Println("Current Time:", now)

	// 🔹 Format time as string
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formatted)

	// 9. STRUCTS AND METHODS
	// 🔹 Define and use a struct
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Name, p.Age) // Output: Alice 25

	// 🔹 Struct with method
	structPerson := Person1{Name: "Alice"}
	structPerson.Greet() // Output: Hello, my name is Alice

	// 🔹 Pass by reference
	x := 10
	changeValue(&x) 
	fmt.Println(x) // Output: 20
}

// changeValue demonstrates pointer usage in Go
func changeValue(num *int) {
	*num = 20
}

// Greet is a method attached to Person1 struct
func (person Person1) Greet() {
	fmt.Println("Hello, my name is", person.Name)
}

// divide demonstrates error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// tulogmoTO demonstrates a goroutine function
func tulogmoTO() {
	fmt.Println("Hello from Goroutine!")
}

// sendMessage sends a message to a channel from a goroutine
func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine!"
}

