package main

import (
	"fmt"
	"strconv"
)

func main() {
	// converting string to int:
	i, err := strconv.Atoi("45")

	// error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)

	}

	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
} else {
		fmt.Println(err)
}

	// simple (short) statement ->  the same effect as the above code
	// i and err are variables scoped to the if statement only

	sample := []interface{}{"apple", "orange"}
	sample = append(sample, "hello")
	fmt.Println(sample[1])

	for _, value := range sample {
		fmt.Println(value)
	}

}
