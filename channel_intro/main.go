package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	c := make(chan int)

	// <- channel operator

	// SEND
	c <- 10

	// RECEIVE
	num := <- c

	fmt.Println(<- c)

	_ = num

	close(c)

	
}