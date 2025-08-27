package beginer

import "fmt"

// ExampleCreateSlice demonstrates creating a slice of structs
// and iterating over nested slices inside the structs.
func ExampleCreateSlice() {
	// Define a struct type Widget with an integer ID
	// and a slice of string attributes
	type Widget struct {
		id    int
		attrs []string
	}

	// Create a slice of Widget instances with initial values
	widgets := []Widget{
		{
			id:    10,
			attrs: []string{"blah", "foo"},
		},
		{
			id:    11,
			attrs: []string{"foo", "bar"},
		},
		{
			id:    12,
			attrs: []string{"xyz"},
		},
	}

	// Iterate over each Widget in the slice
	for _, j := range widgets {
		// Print the Widget ID
		fmt.Printf("%d ", j.id)

		// Iterate over the nested attrs slice and print each attribute
		for _, y := range j.attrs {
			fmt.Printf(" %s ", y)
		}
		// Add a newline after printing each Widget
		fmt.Println()
	}
}
