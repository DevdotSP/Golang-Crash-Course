package beginer

import (
	"fmt"
	"os"
)

// SamplePanic demonstrates how to handle runtime panics using defer and recover
func SamplePanic() {
	// Defer a function to recover from a panic if it occurs
	defer func() {
		if err := recover(); err != nil {
			// Print the panic message to standard error
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			// Exit the program with a non-zero status to indicate error
			os.Exit(1)
		}
	}()

	// Attempt to open a file whose path is provided as the first command-line argument
	file, err := os.Open(os.Args[1])
	if err != nil {
		// Print an error message if the file cannot be opened
		fmt.Println("Could not open file")
	}

	// Print the file object (not the content)
	fmt.Printf("%s", file)
}
