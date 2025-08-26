package main

import "fmt"

func main() {

    // 🔹 Structure of an if–else block in Go:
    //
    // if condition {
    //     // do something if condition is true
    // } else if anotherCondition {
    //     // do something else if the first condition is false but this one is true
    // } else {
    //     // do this if none of the above conditions are true (optional)
    // }

    price, inStock := 100, true

    // Example 1: Simple if
    if price >= 80 { 
        // ✅ Parentheses are NOT required in Go (unlike C/Java).
        fmt.Println("Too Expensive")
    }

    // Example 2: if with logical AND
    if price <= 100 && inStock { 
        // The expression can be shortened to: if price <= 100 && inStock { }
        fmt.Println("Buy it!")
    }

    // ❌ Go does NOT allow "truthy" values like in Python or JavaScript.
    // Example (this would be a compile error):
    // if price {
    //     fmt.Println("We have price!")
    // }

    // Example 3: if–else if–else chain
    // 🔸 Only one branch will be executed, from top to bottom.
    if price < 100 {
        fmt.Println("It's cheap!")
    } else if price == 100 {
        fmt.Println("On the edge")
    } else {
        // This block runs only if all the above conditions are false.
        // The final "else" is optional.
        fmt.Println("It's Expensive!")
    }
}
