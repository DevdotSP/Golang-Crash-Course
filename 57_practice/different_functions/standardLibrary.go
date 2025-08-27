package diffuc

import (
	"fmt"
	"sort"
	"strings"
)

func StandardLib() {
	// =========================
	// Working with strings using the standard library
	// =========================
	greeting := "Hello there friends!"

	// Check if substring exists (case-sensitive)
	fmt.Println("Contains 'hello!':", strings.Contains(greeting, "hello!")) // false, "Hello" != "hello!"

	// Replace all occurrences of a substring
	fmt.Println("Replace 'hello' with 'hi':", strings.ReplaceAll(greeting, "hello", "hi")) // "Hello" remains unchanged (case-sensitive)

	// Convert string to uppercase
	fmt.Println("Uppercase:", strings.ToUpper(greeting)) // "HELLO THERE FRIENDS!"

	// Find the index of a substring (first occurrence)
	fmt.Println("Index of 'th':", strings.Index(greeting, "th")) // returns 6 (position of "th")

	// Split string by a delimiter (here double space, returns single-element slice)
	fmt.Println("Split by double space:", strings.Split(greeting, "  ")) // ["Hello there friends!"]

	fmt.Println("Original string value =", greeting) // strings are immutable

	// =========================
	// Working with slices and sorting
	// =========================
	ages := []int{45, 25, 23, 32, 25, 33, 22, 11}

	// Sort integers in ascending order
	sort.Ints(ages)
	fmt.Println("Sorted ages:", ages) // [11 22 23 25 25 32 33 45]

	// Search for an element in a sorted slice
	index := sort.SearchInts(ages, 90)
	fmt.Println("Index of 90 in sorted ages:", index) // returns len(ages) if not found

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// Sort strings alphabetically
	sort.Strings(names)
	fmt.Println("Sorted names:", names) // [bowser luigi mario peach yoshi]

	// Search for a string in a sorted slice
	fmt.Println("Index of 'bowser':", sort.SearchStrings(names, "bowser")) // returns 0
}
