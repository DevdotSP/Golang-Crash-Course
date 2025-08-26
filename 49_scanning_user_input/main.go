package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a scanner to read input from standard input (keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type something (type 'exit' to quit):")

	// Keep reading user input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of input
		fmt.Println("You entered:", text)

		// If user types "exit", stop the loop and quit
		if text == "exit" {
			fmt.Println("Exiting the scanning ...")
			break
		}
	}

	// Check if the scanner encountered any error while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// This will always run after the loop (when exited normally or via "exit")
	fmt.Println("Just a message after the for loop.")
}
