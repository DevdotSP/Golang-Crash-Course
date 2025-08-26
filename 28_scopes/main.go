// In Go, there are 3 main variable scopes (visibility rules):
//  1. File Scope   -> applies only to imports
//  2. Package Scope -> visible across all files in the same package
//  3. Block Scope   -> visible only inside { } where they are declared

package main

// 🔹 File Scope:
// All import statements are only visible inside THIS file (not across the whole package).
import (
    "fmt"

    // Using "fmt" twice would be an error, since names must be unique in the same scope.
    // But we can import the same package with an alias:
    f "fmt"
)

// 🔹 Package Scope:
// Variables/constants declared outside of any function
// are visible to all functions within the same package.
const done = false

// 🔹 Package-level functions are also package scoped.
func main() {
    // 🔹 Block Scope:
    // This variable is only visible inside the `main` function.
    var task = "Running:"
    fmt.Println(task, done) // Prints: "Running: false" (uses `done` from package scope)

    // Here we use the aliased import `f` instead of `fmt`.
    f.Println("Bye bye!")

    // Names must be unique only within the same scope.
    // Here we declare a new constant `done` inside main().
    // This shadows (overrides) the package-level `done` within this block.
    const done = true
    fmt.Printf("done in main(): %v\n", done) // Prints: true (local `done`)

    f1()
}

func f1() {
    // This function can still "see" the package-scoped `done`
    // because it doesn't declare its own `done`.
    fmt.Printf("done in f1(): %v\n", done) // Prints: false (package-level `done`)
}
