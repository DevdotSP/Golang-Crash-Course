package diffuc

import "fmt"

func Print() {
	age := 35
	name := "Shaun"

	// --- Basic print ---
	fmt.Print("Hello,  ")        // prints without newline
	fmt.Print("world! \n")       // prints with explicit newline
	fmt.Print("new line \n")     // prints with explicit newline

	// --- Println adds a newline automatically ---
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name) // automatically adds spaces between values

	// --- Printf allows formatted strings ---
	fmt.Printf("my age is %v and my name is %v\n", age, name)  // %v = default format
	fmt.Printf("my age is %q and my name is %q\n", age, name)  // %q = quoted strings
	fmt.Printf("age is of type %T\n", age)                     // %T = type
	fmt.Printf("you scored %f points!\n", 255.55)             // %f = float
	fmt.Printf("you scored %0.1f points!\n", 255.25)          // %0.1f = 1 decimal place

	// --- Sprintf saves formatted string into variable ---
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("The saved string is:", str)
}
