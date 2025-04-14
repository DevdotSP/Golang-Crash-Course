package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

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
