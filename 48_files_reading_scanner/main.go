package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file "my_file.txt" for reading
	// If the file does not exist or cannot be opened, the program will log the error and exit
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Always close the file once the function exits (good practice to avoid leaks)
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	// bufio.ScanWords means the scanner will read the file word by word (instead of line by line).
	scanner.Split(bufio.ScanWords)

	// Try scanning the first token (word) in the file
	success := scanner.Scan()

	// If no word was scanned...
	if !success {
		// Check if an error occurred while scanning
		err = scanner.Err()
		if err == nil {
			// If no error, that means we simply reached EOF (empty file or no words to scan)
			log.Println("Scan was completed and it reached EOF")
		} else {
			// If an error occurred during scanning, log and stop the program
			log.Fatal(err)
		}
	}

	// If the first scan was successful, print the first word found
	fmt.Println("First word found:", scanner.Text())

	// Continue scanning through the rest of the file word by word
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// After finishing the scan loop, check if an error occurred during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
