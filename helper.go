package main

import (
	"os"

	"github.com/davecgh/go-spew/spew" // 🔹 Library for dumping Go values in a human-readable format
)


// 1. Essentially, this is a debugging utility:
// 2. spew.Dump(v) = pretty-print a variable
// 3. os.Exit(1) = immediately terminate program
// 4. Dump() = quick "dump and die" like Laravel's dd()



// Dump is a helper function for debugging.
// It pretty-prints any Go value using spew.Dump and then exits the program immediately.
// This is similar to "dd()" in Laravel/PHP (dump and die).
func Dump(v any) {
	spew.Dump(v)   // print value in full detail (structs, pointers, nested fields, etc.)
	os.Exit(1)     // force program to stop after dumping
}

func main() {
	helloWorlds := "hello"

	// Call Dump() to inspect helloWorlds and stop execution.
	// Instead of continuing, it will show "hello" and exit.
	Dump(helloWorlds)
}


