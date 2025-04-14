package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type something (type 'exit' to quit):")

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You entered:", text)

		if text == "exit" {
			fmt.Println("Exiting the scanning ...")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println("Just a message after the for loop.")
}
