package main

import "fmt"

func main() {
	// 1️⃣ Standard for loop (like in C/Java/Python `for`)
	// This loop runs while the condition (i < 10) is true.
	for i := 0; i < 10; i++ {
		fmt.Println("Counting up:", i)
	}

	// 2️⃣ For loop acting like a while loop
	// Go does not have a `while` keyword, but you can achieve the same effect
	// by using `for` with only a condition.
	j := 10
	for j >= 0 {
		fmt.Println("Counting down:", j)
		j--
	}

	// 3️⃣ For loop with multiple variables
	// You can initialize and update multiple variables at once.
	for i, j := 0, 100; i < 5; i, j = i+1, j+10 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	// 4️⃣ Infinite loop
	// A for loop without condition runs forever unless you break manually.
	// sum := 0
	// for {
	//     sum++
	//     if sum > 5 {
	//         break // exit loop when condition met
	//     }
	// }
	// fmt.Println("Sum after break:", sum)

	// 5️⃣ Looping through a slice using `range`
	// Range gives both index and value.
	numbers := []int{10, 20, 30, 40, 50}
	for idx, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", idx, value)
	}

	// 6️⃣ Looping through a string using `range`
	// Range works for Unicode strings (runes).
	word := "GoLang"
	for idx, char := range word {
		fmt.Printf("Character at index %d = %c\n", idx, char)
	}

	// 7️⃣ Looping through a map using `range`
	// Map iteration order is random in Go.
	studentScores := map[string]int{"Alice": 90, "Bob": 85, "Charlie": 92}
	for name, score := range studentScores {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}

	// 8️⃣ Nested for loops (loop inside another loop)
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 2; y++ {
			fmt.Printf("x=%d, y=%d\n", x, y)
		}
	}
}
