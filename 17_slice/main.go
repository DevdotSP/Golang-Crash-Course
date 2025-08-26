package main

import "fmt"

func main() {

	// ============================================================
	// 🔹 Declaring slices
	// ============================================================

	// Declaring a string slice (default = nil, uninitialized)
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil) // -> true
	fmt.Printf("cities: %#v\n", cities)                    // -> []string(nil)

	// ❌ Cannot assign elements to a nil slice (runtime error)
	// cities[0] = "Paris"

	// Declaring and initializing a slice using a slice literal
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers) // -> [2 3 4 5]

	// Creating a slice using make()
	// make([]T, len) → creates slice with given length, filled with zero values
	nums := make([]int, 2) // equivalent to []int{0, 0}
	fmt.Println(nums)      // -> [0 0]

	// Custom slice type
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// ============================================================
	// 🔹 Accessing & modifying slices
	// ============================================================

	// Access element by index
	x := numbers[0]
	fmt.Println("x is", x) // -> x is 2

	// Modify element by index
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // -> []int{2, 200, 4, 5}

	// Iterating with range
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// Iterating with traditional for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])
	}

	// Slices with the same element type can be assigned
	var n []int
	n = numbers
	_ = n

	// ============================================================
	// 🔹 Comparing slices
	// ============================================================

	// Only allowed comparison: slice vs nil
	var nn []int
	fmt.Println(nn == nil) // -> true (nil slice)

	mm := []int{}          // empty but initialized
	fmt.Println(mm == nil) // -> false

	// ❌ Not allowed: direct slice-to-slice comparison
	// fmt.Println(nn == mm)

	// ✅ Compare slices manually (length + elements)
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true

	if len(a) != len(b) {
		eq = false
	}
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	// ============================================================
	// 🔹 Calling helper functions
	// ============================================================
	SliceExpression()
	SliceBackArray()
	AppendingSlices()
}

func SliceExpression() {
	// Arrays, slices, and strings are sliceable
	// Slicing returns a new slice (does NOT copy data)

	a := [5]int{1, 2, 3, 4, 5}

	// Slice expression: a[start:stop] → includes start, excludes stop
	b := a[1:3]
	fmt.Printf("Type: %T , Value: %#v\n", b, b) // -> []int{2, 3}

	// Slice literals
	s1 := []int{1, 2, 3, 4, 5, 6}

	// Example slices
	s2 := s1[1:3]
	fmt.Println(s2) // -> [2 3]

	// Omitting indices
	fmt.Println(s1[2:]) // -> [3 4 5 6]
	fmt.Println(s1[:3]) // -> [1 2 3]
	fmt.Println(s1[:])  // -> [1 2 3 4 5 6]

	// ❌ Out-of-bounds panic
	// fmt.Println(s1[:45])

	// Append to slices
	s1 = append(s1[:4], 100) // replaces everything after index 3 with 100
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200)
	fmt.Println(s1) // -> [1 2 3 4 200]
}

// ============================================================
// 🔹 Slice internals: backing array
// ============================================================
// A slice header contains 3 fields:
// 1. Pointer to backing array
// 2. Length (len())
// 3. Capacity (cap())
// Nil slice → no backing array → all fields are zero
func SliceBackArray() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600 // modifies backing array → all slices reflect change
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]

	// Array as backing array
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() → new slice, can have a different backing array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // copy values, not reference
	cars[0] = "Nissan"
	fmt.Println("cars:", cars, "newCars:", newCars)
	// -> cars: [Nissan Honda Audi Range Rover]
	// -> newCars: [Ford Honda]
}

func AppendingSlices() {
	numbers := []int{2, 3}

	// Append single element
	numbers = append(numbers, 10)
	fmt.Println(numbers) // -> [2 3 10]

	// Append multiple elements
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) // -> [2 3 10 20 30 40]

	// Append all elements of another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... = unpack elements
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	// ============================================================
	// 🔹 Length vs Capacity
	// ============================================================

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // -> 1, 1

	nums = append(nums, 2) // grow slice
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // -> 2, 2

	nums = append(nums, 3) // exceeds capacity → new backing array
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // -> 3, 4

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // -> 5, 8

	// ============================================================
	// 🔹 Copying slices
	// ============================================================

	src := []int{10, 20, 30}
	dst := make([]int, len(src)) // allocate same length
	nn := copy(dst, src)         // copies elements, returns count
	fmt.Println(src, dst, nn)    // -> [10 20 30] [10 20 30] 3
}
