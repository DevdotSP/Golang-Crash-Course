package main

import "fmt"

type Car struct {
	Name  string
	Model string
}

func main() {

	p := Car{"Tricycle", "model"}
	ptr := &p // Pointer to struct

	fmt.Println(*ptr,p)
	fmt.Println("Before:", p.Name) // 25

	ptr.Name = "Twocycle" // Modify via pointer

	fmt.Println("After:", p.Name) // 30

	// x := 300
	// fmt.Printf("Before devalue: %d\n", x)
	// sqp := &x

	// *sqp = 200
	// fmt.Printf("After devalue: %d and %d \n", *sqp, x)
	// ptr := createPointer()
	// fmt.Printf("Value at pointer: %d", *ptr) // 100

	// SampleMarshalling()
	// SampleUnmarshalling()
	// SampleOmitEmpty()
	// SampleJSONtoMap()
	// SampleJSONtoIndent()
	// SampleJSONarrays()
	// SampleFiber()
	// CheckValidJSON()
	// ListOfMethods()

}

func createPointer() *int {
	num := 100
	return &num
}
