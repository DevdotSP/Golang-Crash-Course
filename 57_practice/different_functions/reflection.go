package diffuc

import (
	"fmt"
	"reflect"
)

// SampleReflection demonstrates the usage of various reflection features in Go
func SampleReflection() {
	// Uncomment to run individual reflection examples:
	// ReflectCopy()
	// ReflectDeepEqual()
	// ReflectSwapper()
	// ReflectTypeof()
	// ReflectValueof()
	// NumfieldReflect()
	// ReflectField()
	// ReflectByFieldName()
	// ReflectMakeSlice()
	// MakeMapReflect()
	// ReflectMakeChant()
	ReflectMakeFunction()
}

// --- reflect.MakeFunc() ---
// Dynamically create a function at runtime using reflection
type Sum func(int64, int64) int64

func ReflectMakeFunction() {
	// Define the type of the function
	t := reflect.TypeOf(Sum(nil))

	// Create a function using reflect.MakeFunc
	mul := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		return []reflect.Value{reflect.ValueOf(a + b)}
	})

	// Convert the reflect.Value back to the concrete type
	fn, ok := mul.Interface().(Sum)
	if !ok {
		return
	}

	fmt.Println("Result of dynamically created Sum function:", fn(5, 6))
}

// --- reflect.MakeChan() ---
// Create a new channel dynamically using reflection
func ReflectMakeChant() {
	var str chan string
	strType := reflect.ValueOf(&str)
	newChannel := reflect.MakeChan(reflect.Indirect(strType).Type(), 512)

	fmt.Println("Kind:", newChannel.Kind())       // channel
	fmt.Println("Capacity:", newChannel.Cap())   // 512
}

// --- reflect.MakeMap() ---
// Create a new map dynamically using reflection
func MakeMapReflect() {
	var str map[string]string
	strType := reflect.ValueOf(&str)
	newMap := reflect.MakeMap(reflect.Indirect(strType).Type())

	fmt.Println("Kind:", newMap.Kind()) // map
}

// --- reflect.MakeSlice() ---
// Create a new slice dynamically using reflection
func ReflectMakeSlice() {
	var str []string
	strType := reflect.ValueOf(&str)
	newSlice := reflect.MakeSlice(reflect.Indirect(strType).Type(), 10, 15)

	fmt.Println("Kind:", newSlice.Kind())       // slice
	fmt.Println("Length:", newSlice.Len())     // 10
	fmt.Println("Capacity:", newSlice.Cap())   // 15
}

// --- reflect.FieldByName() ---
// Access and modify struct fields by name dynamically
type TT struct {
	A int
	B string
	C float64
}

func ReflectByFieldName() {
	s := TT{10, "ABCD", 15.20}

	// Access fields dynamically
	fmt.Println("Before modification:")
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))

	// Modify fields dynamically
	reflect.ValueOf(&s).Elem().FieldByName("A").SetInt(50)
	reflect.ValueOf(&s).Elem().FieldByName("B").SetString("Test")
	reflect.ValueOf(&s).Elem().FieldByName("C").SetFloat(5.5)

	fmt.Println("After modification:")
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))
}

// --- reflect.Copy() ---
// Copy elements from one slice to another dynamically
func ReflectCopy() {
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})

	count := reflect.Copy(destination, source)
	fmt.Println("Elements copied:", count)
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}

// --- reflect.DeepEqual() ---
// Compare slices, arrays, and structs for equality
type mobile struct {
	price float64
	color string
}

func ReflectDeepEqual() {
	// Compare slices
	s1 := []string{"A", "B", "C", "D", "E"}
	s2 := []string{"D", "E", "F"}
	fmt.Println("Slices equal?", reflect.DeepEqual(s1, s2))

	// Compare arrays
	n1 := [5]int{1, 2, 3, 4, 5}
	n2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Arrays equal?", reflect.DeepEqual(n1, n2))

	// Compare structs
	m1 := mobile{500.50, "red"}
	m2 := mobile{400.50, "black"}
	fmt.Println("Structs equal?", reflect.DeepEqual(m1, m2))
}

// --- reflect.Swapper() ---
// Create a function to swap elements in a slice dynamically
func ReflectSwapper() {
	theList := []int{1, 2, 3, 4, 5}
	swap := reflect.Swapper(theList)
	fmt.Println("Original slice:", theList)

	swap(1, 3) // Swap elements
	fmt.Println("After swap:", theList)

	// Reverse the slice
	for i := 0; i < len(theList)/2; i++ {
		swap(i, len(theList)-1-i)
	}
	fmt.Println("After reverse:", theList)
}

// --- reflect.TypeOf() ---
// Get the type of variables
func ReflectTypeof() {
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(v1))

	v2 := "Hello World"
	fmt.Println(reflect.TypeOf(v2))

	v3 := 1000
	fmt.Println(reflect.TypeOf(v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println(reflect.TypeOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(v5))

	v6 := true
	fmt.Println(reflect.TypeOf(v6))
}

// --- reflect.ValueOf() ---
// Get the reflect.Value of variables
func ReflectValueof() {
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(v1))

	v2 := "Hello World"
	fmt.Println(reflect.ValueOf(v2))

	v3 := 1000
	fmt.Println(reflect.ValueOf(v3))
	fmt.Println(reflect.ValueOf(&v3)) // pointer value

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println(reflect.ValueOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(v5))

	v6 := true
	fmt.Println(reflect.ValueOf(v6))
}

// --- reflect.NumField() ---
// Count the number of fields in a struct
type T struct {
	A int
	B string
	C float64
	D bool
}

func NumfieldReflect() {
	t := T{10, "ABCD", 15.20, true}
	typeT := reflect.TypeOf(t)
	fmt.Println("Number of fields:", typeT.NumField())
}

// --- reflect.Field() ---
// Iterate struct fields and print names and types
func ReflectField() {
	t := T{A: 10, B: "ABCD", C: 15.20, D: true}
	typeT := reflect.TypeOf(t)

	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println("Field Name:", field.Name, "Field Type:", field.Type)
	}
}

// --- reflect.FieldByIndex() ---
// Access nested fields by index
type ReflectFirst struct {
	A int
	B string
	C float64
}

type reflectSecond struct {
	ReflectFirst
	D bool
}

func ReflectExampleByIndex() {
	s := reflectSecond{ReflectFirst: ReflectFirst{10, "ABCD", 15.20}, D: true}
	t := reflect.TypeOf(s)

	fmt.Printf("FieldByIndex 0: %#v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("FieldByIndex 0,0: %#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("FieldByIndex 0,1: %#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("FieldByIndex 0,2: %#v\n", t.FieldByIndex([]int{0, 2}))
	fmt.Printf("FieldByIndex 1: %#v\n", t.FieldByIndex([]int{1}))
}
