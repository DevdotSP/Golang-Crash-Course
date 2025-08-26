package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func main() {

    // In Go, single quotes represent a character literal, also called a rune.
    // A rune is just an alias for int32 and represents a Unicode code point.
    var1 := 'a'
    fmt.Printf("Type: %T, Value:%d \n", var1, var1) // => Type: int32, Value:97
    // The rune 'a' corresponds to Unicode code point 97.

    // Declaring a string that contains non-ASCII characters
    str := "ţară" // "ţară" means "country" in Romanian.
    // Each character ('ţ', 'a', 'r', 'ă') is a rune.
    // Runes can take between 1 and 4 bytes in UTF-8 encoding.

    // len() returns the number of BYTES, not runes.
    fmt.Println(len(str)) // -> 6 (because "ţ" and "ă" take 2 bytes each, so total 6 bytes)

    // utf8.RuneCountInString() returns the number of runes (characters).
    m := utf8.RuneCountInString(str)
    fmt.Println(m) // -> 4 (ţ, a, r, ă)

    // Accessing by index gives the raw byte value, not the full rune.
    fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163 (second byte of 'ţ')

    // Looping through string using a normal for loop (byte by byte):
    // This prints broken characters, because it treats each byte as a character.
    for i := 0; i < len(str); i++ {
        fmt.Printf("%c", str[i]) // -> Å£arÄ
    }

    fmt.Println("\n" + strings.Repeat("#", 10))

    // Decoding string rune by rune manually using utf8.DecodeRuneInString()
    for i := 0; i < len(str); {
        // DecodeRuneInString reads one rune starting from str[i:]
        r, size := utf8.DecodeRuneInString(str[i:])
        
        // r = rune (actual Unicode character)
        // size = number of bytes that rune takes
        fmt.Printf("%c", r) // -> ţară

        i += size // Move to the next rune by skipping the correct number of bytes
    }

    fmt.Println("\n" + strings.Repeat("#", 10))

    // Easier way: using range automatically decodes runes
    for i, r := range str {
        // i = starting byte index of this rune
        // r = the actual rune (Unicode character)
        fmt.Printf("%d -> %c ", i, r) // -> prints each rune with its byte index
    }
}
