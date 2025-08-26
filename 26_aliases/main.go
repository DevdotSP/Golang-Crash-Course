package main

import "fmt"

func main() {

    // ------------------------------------------------------------
    // 1️⃣ Declaring variables of uint8 and byte types
    // ------------------------------------------------------------
    // 'uint8' is an unsigned 8-bit integer (0 to 255)
    var a uint8 = 10

    // 'byte' is an alias for uint8, used for clarity when working with raw data
    var b byte

    // Even though the variable names differ, 'byte' and 'uint8' are the same type.
    // So assigning a uint8 value to a byte variable does not require conversion.
    b = a
    _ = b // prevent "unused variable" error

    // ------------------------------------------------------------
    // 2️⃣ Creating a type alias
    // ------------------------------------------------------------
    // 'type second = uint' creates an alias named 'second' for the type 'uint'
    // Aliases allow using alternative names for existing types without creating a new type
    type second = uint

    // Using the alias
    var hour second = 3600
    fmt.Printf("hour type: %T\n", hour) // => hour type: uint

    // Since 'second' is an alias for uint, arithmetic operations don't require conversion
    fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60

    // ------------------------------------------------------------
    // Notes:
    // - Aliases differ from 'new types' created with 'type newType originalType',
    //   which are distinct types requiring explicit conversion.
    // - byte and rune are built-in aliases for uint8 and int32, respectively.
    // ------------------------------------------------------------
}
