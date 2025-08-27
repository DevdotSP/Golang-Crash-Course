package diffuc

import (
	"encoding/json"
	"fmt"
)

func ExampleStruct() {
	// Uncomment to test examples
	// BasicStruct()                   // Simple struct creation
	// CreateInstanceUsingDot()        // Assign values using dot notation
	// WaysToAssignValuesToStruct()    // Different ways to initialize struct
	// UsingNewKeyword()               // Using `new` to create struct pointer
	// UsingPointer()                  // Working with struct pointers
	// NestedStruct()                  // Struct containing slices and nested structs
	// UsingJsonTag()                  // JSON marshal/unmarshal using struct tags
	// CopyStructUsingValuePointer()   // Copying structs by value and pointer
	AssignDefaultValueOnStruct()       // Assign default values when struct fields are zero
}

// ----------------------
// Basic struct example
// ----------------------
func BasicStruct() {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	// Print a struct literal directly
	fmt.Println(rectangle{10.5, 25.10, "red"})
}

// ----------------------
// Assign values using dot notation
// ----------------------
func CreateInstanceUsingDot() {
	type rectangle struct {
		length  int
		breadth int
		color   string

		geometry struct {
			area      int
			perimeter int
		}
	}

	var rect rectangle
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	// Compute derived values in nested struct
	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}

// ----------------------
// Different ways to assign values to struct fields
// ----------------------
func WaysToAssignValuesToStruct() {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	// Assign all values in order
	rect1 := rectangle{10, 20, "Green"}
	fmt.Println(rect1)

	// Assign selected fields by name, others default to zero
	rect2 := rectangle{length: 10, color: "Green"}
	fmt.Println(rect2)

	// Shorthand initialization
	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)

	// Named fields initialization
	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	// Partial fields, omitted fields default to zero
	rect5 := rectangle{breadth: 20, color: "Green"}
	fmt.Println(rect5)
}

// ----------------------
// Using `new` keyword to create struct pointer
// ----------------------
func UsingNewKeyword() {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	rect1 := new(rectangle) // rect1 is a pointer to rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	rect2 := new(rectangle)
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}

// ----------------------
// Working with pointers to structs
// ----------------------
func UsingPointer() {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	rect1 := &rectangle{10, 20, "Green"} // Must provide all fields
	fmt.Println(rect1)

	rect2 := &rectangle{} // Can assign fields individually
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)

	rect3 := &rectangle{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3)
}

// ----------------------
// Nested structs example
// ----------------------
func NestedStruct() {
	type Salary struct {
		Basic, HRA, TA float64
	}

	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{Basic: 15000, HRA: 5000, TA: 2000},
			{Basic: 16000, HRA: 5000, TA: 2100},
			{Basic: 17000, HRA: 5000, TA: 2200},
		},
	}

	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for i, s := range e.MonthlySalary {
		fmt.Printf("Month %d: %+v\n", i+1, s)
	}
}

// ----------------------
// Using JSON struct tags
// ----------------------
func UsingJsonTag() {
	type Employee struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		City      string `json:"city"`
	}

	// Unmarshal JSON into struct
	jsonStr := `{"firstname": "Rocky", "lastname": "Sting", "city": "London"}`
	emp := new(Employee)
	json.Unmarshal([]byte(jsonStr), emp)
	fmt.Println(emp)

	// Marshal struct into JSON
	emp2 := &Employee{FirstName: "Ramesh", LastName: "Soni", City: "Mumbai"}
	out, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", out)
}

// ----------------------
// Copy struct using value and pointer
// ----------------------
func CopyStructUsingValuePointer() {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	r1 := rectangle{10, 20, "Green"}
	fmt.Println("Original:", r1)

	r2 := r1          // Copy by value (changes don't affect r1)
	r2.color = "Pink"
	fmt.Println("Modified Copy:", r2)

	r3 := &r1         // Copy by pointer (changes affect original)
	r3.color = "Red"
	fmt.Println("Pointer Modified:", r3)

	fmt.Println("Original after pointer modification:", r1)
}

// ----------------------
// Assign default values when struct fields are zero
// ----------------------
type Employee struct {
	Name string
	Age  int
}

func AssignDefaultValueOnStruct() {
	emp1 := Employee{Name: "Mr. Fred"} // Age will be assigned default
	emp1.Info()
	fmt.Println(emp1)

	emp2 := Employee{Age: 26} // Name will be assigned default
	emp2.Info()
	fmt.Println(emp2)

	emp3 := Employee{} // Both fields will get defaults
	emp3.Info()
	fmt.Println(emp3)
}

// Helper method to assign defaults
func (obj *Employee) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}
