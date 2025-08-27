package StringManipulation

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Entry function to call all string manipulation examples
func StringManipulate() {
	UTF8Encoding()      // Demonstrates rune and byte representation
	StringsStatic()     // Count occurrences of a substring
	StringPadding()     // Pad a string using formatting
	StringTrim()        // Trim spaces from a string
	StringTraverse()    // Traverse string using range over runes
	GetCharsInString()  // Access specific byte in a string
	StringJoin()        // Join slice of strings into one string
	StringSplit()       // Split string into a slice
	StringEncrypt()     // MD5 hash encryption example
	RegexProcess()      // Extract substring using regular expressions
	StringConversion()  // Convert string to upper/lower case
	StringReplacement() // Replace characters/substrings
	StringComparison()  // Compare strings for equality
	StringSearch()      // Search substring in a string
	StringConcat()      // Concatenate strings
	ExampleImmutability() // Demonstrates string immutability
	StringSlice()       // Extract substring via slicing
}

// UTF8Encoding shows rune and byte-level representation of a string
func UTF8Encoding() {
	str := "Go世界" // Contains ASCII and Unicode characters

	fmt.Println("String:", str)

	// Print Unicode code points (runes)
	fmt.Println("Rune Code Points:")
	for _, runeValue := range str {
		fmt.Printf("%c: %U\n", runeValue, runeValue)
	}

	// Print UTF-8 encoded bytes
	fmt.Println("UTF-8 Bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c: %x\n", str[i], str[i])
	}
}

// Count occurrences of a substring
func StringsStatic() {
	str := "Go is easy to learn. Go is powerful."
	count := strings.Count(str, "Go")
	fmt.Println("Count of 'Go':", count)
}

// Pad a string (right-aligned example)
func StringPadding() {
	str := "Go"
	padded := fmt.Sprintf("%-10s", str) // Pads with spaces to 10 characters
	fmt.Printf("Padded: '%s'\n", padded)
}

// Trim spaces from start and end of a string
func StringTrim() {
	str := "   Go Lang   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed:", trimmed)
}

// Traverse string by characters (handles Unicode correctly)
func StringTraverse() {
	str := "Go"
	for index, char := range str {
		fmt.Printf("At index %d, char: %c\n", index, char)
	}
}

// Access specific byte in a string (note: may split multi-byte Unicode chars)
func GetCharsInString() {
	str := "Go"
	byteValue := str[1]
	fmt.Println("Byte at index 1:", byteValue)
}

// Join a slice of strings into one string
func StringJoin() {
	items := []string{"apple", "banana", "cherry"}
	str := strings.Join(items, ", ")
	fmt.Println("Joined string:", str)
}

// Split string into a slice based on separator
func StringSplit() {
	str := "apple,banana,cherry"
	items := strings.Split(str, ",")
	fmt.Println("Split string:", items)
}

// MD5 hash of a string
func StringEncrypt() {
	str := "secret data"
	hasher := md5.New()
	io.WriteString(hasher, str)
	fmt.Printf("MD5 Hash: %x\n", hasher.Sum(nil))
}

// Extract substring using regex
func RegexProcess() {
	str := "My email is example@example.com"
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`)
	email := re.FindString(str)
	fmt.Println("Extracted email:", email)
}

// Convert string to lower and upper case
func StringConversion() {
	str := "GoLang"
	lowercase := strings.ToLower(str)
	uppercase := strings.ToUpper(str)
	fmt.Println("Lowercase:", lowercase)
	fmt.Println("Uppercase:", uppercase)
}

// Replace characters or substrings
func StringReplacement() {
	s := "banana"
	r1 := strings.Replace(s, "a", "o", 2) // First 2 occurrences
	r2 := strings.ReplaceAll(s, "a", "o")  // All occurrences
	fmt.Println("Replace first 2:", r1)
	fmt.Println("Replace all:", r2)
}

// Compare strings
func StringComparison() {
	str1 := "hello"
	str2 := "world"
	str3 := "hello"

	fmt.Println("str1 == str2?", str1 == str2)
	fmt.Println("str1 == str3?", str1 == str3)
}

// Search substring in string
func StringSearch() {
	haystack := "Hello, Golang World!"
	needle := "Golang"

	if strings.Contains(haystack, needle) {
		fmt.Println("Found:", needle)
	} else {
		fmt.Println(needle, "not found")
	}
}

// Concatenate strings using +
func StringConcat() {
	str1 := "Hello, "
	str2 := "world!"
	result := str1 + str2
	fmt.Println("Concatenated:", result)
}

// Demonstrate string immutability
func ExampleImmutability() {
	original := "hello"
	modified := original
	modified = "world"

	fmt.Println("Original:", original)
	fmt.Println("Modified:", modified)
}

// Extract substring via slicing
func StringSlice() {
	str := "Hello, World!"
	substring := str[7:12] // from index 7 to 11
	fmt.Println("Substring:", substring)
}
