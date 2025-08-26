package main

import "fmt"

// ✅ Creating new types in Go

// age is a new type based on int
type age int        

// oldAge is a new type based on age, but its underlying type is still int
type oldAge age     

// veryOldAge is another new type based on age, underlying type is still int
type veryOldAge age 

func main() {

    // 🔹 Define a new type speed with underlying type uint
    type speed uint

    // Declaring variables of type speed
    var s1 speed = 10
    var s2 speed = 20

    // Performing arithmetic operations between variables of the same new type
    fmt.Println(s2 - s1) // -> 10

    // ❌ Important: New types are distinct from their underlying types
    var x uint
    // x = s1  // ERROR: cannot use s1 (type speed) as type uint

    // ✅ Correct way: Explicit conversion to the underlying type
    x = uint(s1)
    _ = x

    // ✅ Converting back from underlying type to new type
    var s3 speed = speed(x)
    _ = s3
}
