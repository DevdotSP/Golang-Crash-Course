// =============================================
// VARIABLES IN GO
// =============================================
//
// Go is a statically typed language, meaning the type of
// each variable is known at compile time. There are several
// ways to declare and initialize variables:
//
// 1. Standard Declaration:
//      var name type
//      var x int = 10
//
// 2. Type Inference:
//      var y = 20      // Go infers int
//
// 3. Short Declaration (inside functions only):
//      z := 30
//
// 4. Multiple Declarations:
//      var a, b int
//      var (
//          age float64
//          name string
//          active bool
//      )
//
// Other important notes:
// - Variables must be used (otherwise compile-time error).
// - The Blank Identifier `_` discards unused values.
// - Multiple assignment allows swapping values directly.
// - Strong typing means variables cannot change type.
//
// =============================================

package main

import "fmt"

func main() {

    // ================================
    // 1. DECLARING VARIABLES
    // ================================
    var s1 string       // explicit type declaration
    s1 = "Learning Go!" // assign value later
    fmt.Println(s1)     // Output: Learning Go!

    // ================================
    // 2. TYPE INFERENCE
    // ================================
    // Go automatically deduces the type based on the initial value.
    var k int = 6 // Explicit, but type (int) could be omitted.
    var i = 5     // type inferred as int
    var j = 5.6   // type inferred as float64

    fmt.Println("i:", i, "j:", j, "k:", k)

    // Invalid: can't assign float to int (strongly typed)
    // i = j  // ERROR

    // New variable with type inference
    var s2 = "Go!"
    _ = s2 // use blank identifier to avoid unused variable error

    // ================================
    // 3. MULTIPLE ASSIGNMENT
    // ================================
    var ii, jj int
    ii, jj = 5, 8 // tuple assignment (assign multiple values at once)

    // Swap variables easily using multiple assignment
    ii, jj = jj, ii
    fmt.Println(ii, jj)

    // ================================
    // 4. SHORT DECLARATION ( := )
    // ================================
    // Works only inside functions (block scope).
    // At least one NEW variable must be on the left side.
    s3 := "Learning golang!"
    _ = s3

    // Short declaration with multiple variables
    car, cost := "Audi", 50000
    fmt.Println(car, cost)

    // Redeclaration: at least one new variable required
    var deleted = false
    deleted, file := true, "a.txt" // "file" is new, so this works
    _, _ = deleted, file

    // Expressions allowed in short declarations
    sum := 5 + 2.5
    fmt.Println(sum)

    // ================================
    // 5. MULTIPLE DECLARATIONS
    // ================================
    // Better readability with grouped declarations
    var (
        age       float64
        firstName string
        gender    bool
    )
    _, _, _ = age, firstName, gender

    // Concise multiple declaration of same type
    var a, b int
    _, _ = a, b
}
