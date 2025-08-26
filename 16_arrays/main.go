package main

import (
    "fmt"
)

func main() {

    // ------------------------------------------------------------
    // 1️⃣ Declaring arrays
    // ------------------------------------------------------------
    // An array of 4 integers (default zero-values)
    var numbers [4]int

    // printing array values
    fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
    fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}

    // Array literals
    var a1 = [4]float64{}                           // all elements initialized to 0.0
    var a2 = [3]int{5, -3, 5}                       // initialized with explicit values
    a3 := [4]string{"Dan", "Diana", "Paul", "John"} // short declaration
    a4 := [4]string{"x", "y"}                       // first 2 elements initialized, rest default ""

    // Ellipsis (...) automatically determines the length
    a5 := [...]int{1, 4, 5}
    fmt.Println("The length of a5 is: ", len(a5)) // -> 3

    // Multiline array initialization
    a6 := [...]int{
        1,
        2,
        3,
    } // ending comma is mandatory

    _, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

    // ------------------------------------------------------------
    // 2️⃣ Modifying array elements
    // ------------------------------------------------------------
    numbers[0] = 7              // change first element
    fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

    // Index out-of-bounds would cause compile-time error
    // numbers[5] = 8 // invalid

    // Accessing elements
    x := numbers[0]
    fmt.Println("x is ", x) // -> x is 7

    // ------------------------------------------------------------
    // 3️⃣ Iterating over arrays
    // ------------------------------------------------------------
    // Range-based iteration
    for i, v := range numbers {
        fmt.Println("index:", i, "value:", v)
    }

    // Classic for-loop iteration
    for i := 0; i < len(numbers); i++ {
        fmt.Println("index:", i, "value:", numbers[i])
    }

    // ------------------------------------------------------------
    // 4️⃣ Multi-dimensional arrays (matrix)
    // ------------------------------------------------------------
    balances := [2][3]int{
        {5, 6, 7},
        {8, 9, 10},
    }

    for _, arr := range balances {
        for _, value := range arr {
            fmt.Printf("%d ", value)
        }
        fmt.Println()
    }

    // ------------------------------------------------------------
    // 5️⃣ Copying arrays
    // ------------------------------------------------------------
    m := [3]int{1, 2, 3}
    n := m // n is a copy, not a reference
    fmt.Println("n is equal to m:", n == m) // -> true

    m[0] = -1 // modify only m
    fmt.Println("n is equal to m:", n == m) // -> false

    // ------------------------------------------------------------
    // 6️⃣ Arrays with keyed elements
    // ------------------------------------------------------------
    grades := [3]int{
        1: 10,
        0: 5,
        2: 7,
    }
    fmt.Println(grades) // -> [5 10 7]

    accounts := [3]int{
        2: 50,
    }
    fmt.Println(accounts) // -> [0 0 50]

    names := [...]string{
        4: "Dan",
    }
    fmt.Println(len(names)) // -> 5

    // Unkeyed elements get the next index after the last keyed element
    cities := [...]string{
        5: "Paris",
        "London", // index 6
        1: "NYC", // explicitly keyed
    }
    fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

    // Boolean array example
    weekend := [7]bool{5: true, 6: true}
    fmt.Println(weekend) // -> [false false false false false true true]
}
