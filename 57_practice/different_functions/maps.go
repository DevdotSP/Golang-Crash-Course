package diffuc

import "fmt"

// ExampleMaps demonstrates common map operations in Go
func ExampleMaps() {
	// Uncomment the examples you want to run
	// ExampleDeclaraction()
	// MapDeclarationMakeFunction()
	// MapLength()
	CRUDMap()      // Demonstrates Create, Read, Update, Delete
	// IterateMap() // Demonstrates iteration over a map
}

// ExampleDeclaraction demonstrates map declaration using literals and basic access
func ExampleDeclaraction() {
	// Declare a map with string keys and int values
	var employee = map[string]int{
		"light":   2,
		"raphael": 3,
	}

	// Access value for a key with existence check
	value, exists := employee["raphael"]
	if exists {
		fmt.Println("Value for raphael:", value) // Output: 3
	} else {
		fmt.Println("Key not found")
	}

	value, exists = employee["light"]
	if exists {
		fmt.Println("Value for light:", value) // Output: 2
	} else {
		fmt.Println("Key not found")
	}

	// Print full map and type
	fmt.Println("Full map:", employee)        // map[light:2 raphael:3]
	fmt.Printf("Type of map: %T\n", employee) // map[string]int
}

// MapDeclarationMakeFunction demonstrates creating a map using the make() function
func MapDeclarationMakeFunction() {
	// Create empty map using make
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println("Employee map:", employee)

	// Another example
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println("Employee list map:", employeeList)
}

// MapLength shows how to get the number of elements in a map
func MapLength() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20

	// Empty map
	employeeList := make(map[string]int)

	fmt.Println("Length of employee map:", len(employee))     // 2
	fmt.Println("Length of empty map:", len(employeeList))    // 0
}

// CRUDMap demonstrates Create, Read, Update, Delete operations
func CRUDMap() {
	// Read element
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println("Read Mark:", employee["Mark"])

	// --- Create / Add elements ---
	var employeeADD = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println("Initial Map:", employeeADD)

	employeeADD["Rocky"] = 30
	employeeADD["Josef"] = 40
	fmt.Println("After Adding:", employeeADD)

	// --- Update element ---
	var employeeUPDATE = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println("Initial Map:", employeeUPDATE)
	employeeUPDATE["Mark"] = 50
	fmt.Println("After Update:", employeeUPDATE)

	// --- Delete element ---
	var employeeDELETE = make(map[string]int)
	employeeDELETE["Mark"] = 10
	employeeDELETE["Sandy"] = 20
	employeeDELETE["Rocky"] = 30
	employeeDELETE["Josef"] = 40
	fmt.Println("Before Delete:", employeeDELETE)

	delete(employeeDELETE, "Mark")
	fmt.Println("After Delete:", employeeDELETE)
}

// IterateMap demonstrates iterating over map keys and values
func IterateMap() {
	var employee = map[string]int{
		"Mark":  10,
		"Sandy": 20,
		"Rocky": 30,
		"Rajiv": 40,
		"Kate":  50,
	}

	for key, element := range employee {
		fmt.Printf("Key: %-6s => Element: %d\n", key, element)
	}
}
