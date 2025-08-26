package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ============================================================
	// 1️⃣ Converting between numeric types (int, float64)
	// ============================================================

	var x = 3   // int type
	var y = 3.2 // float64 type

	// x = x * y // ❌ compile error: mismatched types (int * float64)

	// Convert float64 to int before multiplication
	x = x * int(y)
	fmt.Println(x) // => 9

	// Convert int to float64 before multiplication
	y = float64(x) * y
	fmt.Println(y) // => 28.8

	// Convert float64 to int for assignment to int variable
	x = int(float64(x) * y)
	fmt.Println(x) // => 259

	// ============================================================
	// 2️⃣ Types with different names are considered different
	// ============================================================

	var a int = 5    // int (platform-dependent size)
	var b int64 = 2  // int64
	// a = b // ❌ cannot assign int64 to int directly
	a = int(b) // ✅ explicit conversion required

	_ = a // prevent unused variable error

	// ============================================================
	// 3️⃣ Converting numbers to strings and vice versa
	// ============================================================

	// int to string using rune (Unicode code point)
	s := string(99)           
	fmt.Println(s)            // => "c" (ASCII code 99)
	fmt.Println(string(34234)) // => "薺" (Unicode code point 34234)

	// ❌ Cannot convert float directly to string like int
	// s1 := string(65.1) // error

	// Convert float to string using fmt.Sprintf
	myStr := fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr) // => "5.120000"

	// Convert int to string using fmt.Sprintf
	myStr1 := fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // => "34234"

	// Convert string to float64 using strconv.ParseFloat
	result, err := strconv.ParseFloat("3.142", 64)
	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
		fmt.Println("Cannot convert to float64!")
	}

	// ============================================================
	// 4️⃣ Convert string to int and int to string
	// ============================================================

	// string to int using strconv.Atoi
	i, err := strconv.Atoi("-50")
	if err != nil {
		fmt.Println("Cannot convert to int!")
	}

	// int to string using strconv.Itoa
	s = strconv.Itoa(20)

	fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50
	fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string, s value is "20"
}
