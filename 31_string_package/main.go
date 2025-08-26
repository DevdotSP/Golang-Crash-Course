package main

import (
    "fmt"
    "strings"
)

func main() {

    // Assign fmt.Println to a variable for shorthand use
    p := fmt.Println

    // -------------------------------------------
    // Contains() → checks if substring exists in a string
    result := strings.Contains("I love Go Programming!", "love")
    p(result) // -> true

    // ContainsAny() → checks if ANY of the chars in the second string are present in the first
    result = strings.ContainsAny("success", "xy")
    p(result) // -> false

    // ContainsRune() → checks if a specific rune (character) exists in a string
    result = strings.ContainsRune("golang", 'g')
    p(result) // -> true

    // -------------------------------------------
    // Count() → counts the number of non-overlapping occurrences of a substring
    n := strings.Count("cheese", "e")
    p(n) // -> 3

    // Special case: empty substring → returns 1 + number of runes in the string
    n = strings.Count("five", "")
    p(n) // -> 5 (1 + 4 runes)

    // -------------------------------------------
    // ToLower() and ToUpper() → convert case
    p(strings.ToLower("Go Python Java")) // -> go python java
    p(strings.ToUpper("Go Python Java")) // -> GO PYTHON JAVA

    // -------------------------------------------
    // Comparing strings (case-sensitive)
    p("go" == "go") // -> true
    p("Go" == "go") // -> false

    // Case-insensitive comparison (inefficient way: convert both to same case)
    p(strings.ToLower("Go") == strings.ToLower("go")) // -> true

    // EqualFold() → efficient case-insensitive comparison
    p(strings.EqualFold("Go", "gO")) // -> true

    // -------------------------------------------
    // Repeat() → repeats a string n times
    myStr := strings.Repeat("ab", 10)
    p(myStr) // -> abababababababababab

    // Replace() → replaces substrings (3rd arg = number of replacements)
    myStr = strings.Replace("192.168.0.1", ".", ":", 2)
    p(myStr) // -> 192:168:0.1

    // Replace with -1 → replace ALL occurrences
    myStr = strings.Replace("192.168.0.1", ".", ":", -1)
    p(myStr) // -> 192:168:0:1

    // ReplaceAll() → shorthand for replacing all occurrences
    myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
    p(myStr) // -> 192:168:0:1

    // -------------------------------------------
    // Split() → splits a string into a slice by a separator
    s := strings.Split("a,b,c", ",")
    fmt.Printf("%T\n", s)                  // -> []string
    fmt.Printf("strings.Split():%#v\n", s) // -> []string{"a", "b", "c"}

    // Special case: empty separator → split after every rune
    s = strings.Split("Go for Go!", "")
    fmt.Printf("strings.Split():%#v\n", s)
    // -> []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

    // -------------------------------------------
    // Join() → joins slice of strings with a separator
    s = []string{"I", "learn", "Golang"}
    j := strings.Join(s, "-")
    fmt.Printf("%T\n", j) // -> string
    p(j)                  // -> I-learn-Golang

    // -------------------------------------------
    // Fields() → splits string by any whitespace (spaces, tabs, newlines)
    myStr = "Orange Green \n Blue Yellow"
    fields := strings.Fields(myStr)
    fmt.Printf("%T\n", fields)  // -> []string
    fmt.Printf("%#v\n", fields) // -> []string{"Orange", "Green", "Blue", "Yellow"}

    // -------------------------------------------
    // TrimSpace() → removes leading/trailing whitespace (tabs, spaces, newlines)
    s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
    fmt.Printf("%q\n", s1) // "Goodbye Windows, Welcome Linux!"

    // Trim() → removes specified leading/trailing characters
    s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
    fmt.Printf("%q\n", s2) // "Hello, Gophers"
}
