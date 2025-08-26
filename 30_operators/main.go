package main

import "fmt"

func main() {
    a, b := 10, 5.5

    // ---------------------------
    // ARITHMETIC OPERATORS
    // ---------------------------
    // +   addition
    // -   subtraction
    // *   multiplication
    // /   division (integer division if both operands are int)
    // %   modulus (remainder after integer division)
    // NOTE: Go does not have an exponentiation operator.
    //       Use math.Pow(float64, float64) for powers.
    fmt.Println(a + 5)   // 10 + 5 = 15
    fmt.Println(3.1 - b) // 3.1 - 5.5 = -2.4
    fmt.Println(a * a)   // 10 * 10 = 100
    fmt.Println(a / a)   // 10 / 10 = 1
    fmt.Println(11 / 5)  // Integer division → 2 (remainder discarded)

    // Go is a strongly typed language: you cannot mix int and float64 directly.
    // fmt.Println(a * b) // ERROR: mismatched types int and float64
    fmt.Println(a * int(b))     // Convert b to int → 10 * 5 = 50
    fmt.Println(float64(a) * b) // Convert a to float64 → 10.0 * 5.5 = 55.0

    // ---------------------------
    // INCREMENT / DECREMENT
    // ---------------------------
    // ++ increases a variable by 1
    // -- decreases a variable by 1
    // (⚠️ only valid as a statement, not in expressions like x = y++)
    x := 10
    x++ // x becomes 11 (same as: x += 1)
    x-- // x becomes 10 (same as: x -= 1)

    // ---------------------------
    // ASSIGNMENT OPERATORS
    // ---------------------------
    // =   assign
    // +=  add and assign
    // -=  subtract and assign
    // *=  multiply and assign
    // /=  divide and assign
    // %=  modulus and assign
    a = 10
    a += 2 // a = 10 + 2 → 12
    a -= 3 // a = 12 - 3 → 9
    a *= 2 // a = 9 * 2  → 18
    a /= 3 // a = 18 / 3 → 6
    a %= 5 // a = 6 % 5  → 1

    // ---------------------------
    // COMPARISON OPERATORS
    // ---------------------------
    // ==  equal to
    // !=  not equal to
    // >   greater than
    // <   less than
    // >=  greater than or equal
    // <=  less than or equal
    fmt.Println(5 == 6)   // false
    fmt.Println(5 != 6)   // true
    fmt.Println(10 > 10)  // false
    fmt.Println(10 >= 10) // true
    fmt.Println(5 < 5)    // false
    fmt.Println(5 <= 5)   // true

    // ---------------------------
    // LOGICAL OPERATORS
    // ---------------------------
    // &&  logical AND (true if both operands are true)
    // ||  logical OR  (true if at least one operand is true)
    // !   logical NOT (inverts a boolean value)
    fmt.Println(0 < 2 && 4 > 1) // true  (both conditions true)
    fmt.Println(1 > 5 || 4 > 5) // false (both conditions false)
    fmt.Println(!(1 > 2))       // true  (since 1 > 2 is false, !false = true)
}
