package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
Step-by-step explanation:

1. Open and Read Specific Bytes (io.ReadFull)
   - Opens test.txt.
   - Reads exactly 2 bytes into byteSlice.
   - If the file has fewer than 2 bytes, it returns an error.
   - Useful when you need a fixed-size chunk (e.g., reading headers or magic numbers in binary files).

2. Read Entire File into Memory (io.ReadAll)
   - Opens main.go.
   - Reads the entire file contents into memory as a []byte.
   - Converts it into a string for printing.
   - Reports the number of bytes read.

3. Read Another File with io.ReadAll
   - Repeats the process with test.txt.
   - Demonstrates reusing the same logic for multiple files.
*/

func main() {
	// === 1. Read fixed-size chunk with io.ReadFull ===
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2) // prepare a 2-byte buffer

	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	// === 2. Read entire file with io.ReadAll ===
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	// === 3. Re-read another file with io.ReadAll ===
	file, err = os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", data)
}
