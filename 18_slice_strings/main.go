package main

import "fmt"

func main() {

    // 🔹 Example 1: Slicing a string with ASCII characters
    // In Go, a string is a sequence of bytes.
    // Slicing a string reuses the same backing array (no copy happens).
    // But the slice result gives raw bytes, not Unicode runes.
    s1 := "abcdefghijkl"
    fmt.Println(s1[2:5]) // -> "cde", takes bytes from index 2 (inclusive) to 5 (exclusive).

    // 🔹 Example 2: Slicing a string with Unicode characters (UTF-8 encoded).
    // Each Chinese character takes 3 bytes in UTF-8.
    // Direct slicing gives raw bytes, not full characters.
    // This may result in invalid characters (�).
    s2 := "中文维基是世界上"
    fmt.Println(s2[0:2]) // -> "�" because only the first 2 bytes of "中" are taken, not the full 3-byte character.

    // 🔹 Correct way: Convert string to a slice of runes.
    // A rune in Go represents a Unicode code point (int32).
    rs := []rune(s2)
    fmt.Printf("%T\n", rs) // -> []int32, each element represents a full character.

    // 🔹 Now slicing works safely on runes (not bytes).
    // Here we slice the first 3 characters (not bytes).
    fmt.Println(string(rs[0:3])) // -> "中文维"
}
