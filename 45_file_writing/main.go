package main

import (
	"log"
	"os"
)

func main() {

	// ✅ OpenFile explanation:
	// os.OpenFile lets us open (or create) a file with custom flags and permissions.
	// In this example:
	// - os.O_WRONLY → open the file in "write-only" mode
	// - os.O_TRUNC  → if the file already exists, clear (truncate) its content before writing
	// - os.O_CREATE → if the file does not exist, create a new one with the given permissions
	//
	// Permissions: 0644 means:
	//   - owner: read + write
	//   - group: read only
	//   - others: read only
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		// If file opening/creation fails, stop the program and log the error
		log.Fatal(err)
	}
	// Always close the file when you're done
	defer file.Close()

	// -------------------------------------------------------
	// 1️⃣ WRITING BYTES TO FILE (manual approach)
	// -------------------------------------------------------

	// Convert string into a slice of bytes
	byteSlice := []byte("I learn Golang! 传")

	// Write the byte slice into the file
	// Returns: (number of bytes written, error)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	// Print how many bytes were written
	log.Printf("Bytes written: %d\n", bytesWritten)
	// Example output: Bytes written: 19

	// -------------------------------------------------------
	// 2️⃣ WRITING BYTES TO FILE (simplified approach)
	// -------------------------------------------------------

	// os.WriteFile() is a convenience function.
	// It does everything in one step:
	//   - creates the file (if it doesn’t exist)
	//   - opens it
	//   - truncates it if it already exists
	//   - writes the content
	//   - closes the file
	//
	// (Note: ioutil.WriteFile() was used before Go 1.16, now replaced by os.WriteFile)

	bs := []byte("Go Programming is cool!")
	err = os.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
