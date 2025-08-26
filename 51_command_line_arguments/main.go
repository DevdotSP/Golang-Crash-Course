package main

import (
	"fmt"
	"os"
)

func main() {
	// ============================================================
	// 1️⃣ os.Args Overview
	// ============================================================
	// os.Args is a slice of strings ([]string) that stores
	// the command-line arguments passed to the program.
	// The first element, os.Args[0], is always the program's path.
	// Any additional elements are the arguments passed to the program.

	fmt.Println("os.Args:", os.Args) // print the whole slice of arguments

	// ============================================================
	// 2️⃣ Accessing individual command-line arguments
	// ============================================================

	// os.Args[0] -> program name/path
	fmt.Println("Path:", os.Args[0])

	// os.Args[1] -> first argument passed to the program
	// Make sure to check len(os.Args) before accessing to avoid index out of range errors
	if len(os.Args) > 1 {
		fmt.Println("1st Argument:", os.Args[1])
	}

	// os.Args[2] -> second argument passed to the program
	if len(os.Args) > 2 {
		fmt.Println("2nd Argument:", os.Args[2])
	}

	// ============================================================
	// 3️⃣ Checking the number of arguments
	// ============================================================
	fmt.Println("No. of items inside os.Args:", len(os.Args))
}
