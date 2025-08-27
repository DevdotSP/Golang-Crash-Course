package beginer

import "fmt"

// Before is a simple struct with a single string field
type Before struct {
	m string
}

// append demonstrates using empty interfaces to return a new struct
// It takes any type as input (here we expect Before) and returns
// a new struct that embeds the original struct and adds a new field `n`.
func append(b interface{}) interface{} {
	return struct {
		Before // embeds the original struct
		n      string
	}{
		b.(Before), // type assertion to extract the Before struct
		"rest",     // new field value
	}
}

// SampleEmptyInterface shows how to use the append function
// with empty interfaces to create a new struct dynamically
func SampleEmptyInterface() {
	b := Before{"test"}  // create an instance of Before
	a := append(b)       // call append function to create a new struct
	fmt.Println(a)       // print the resulting struct
}
