package main

import "fmt"

func main() {

	// ------------------------------------
	// DECLARING A MAP
	// ------------------------------------
	// A map is a built-in data type that associates keys with values.
	// Syntax: map[KeyType]ValueType
	var employees map[string]string // nil map (not initialized yet)

	fmt.Printf("%#v\n", employees) // -> map[string]string(nil)
	fmt.Printf("No. of elements: %d\n", len(employees)) // => 0

	// Accessing a key that doesn't exist (or on a nil map)
	// returns the *zero value* for the value type.
	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"]) // => ""

	// NOTE: Map keys must be of comparable types (== or != must work).
	// Example: slices cannot be keys → invalid.
	// var m1 map[[]int]string // ❌ ERROR

	// ❌ You cannot insert into a nil map (will panic):
	// employees["Dan"] = "Programmer"

	// ------------------------------------
	// INITIALIZING MAPS
	// ------------------------------------

	// ✅ Using a map literal (empty map)
	people := map[string]float64{}
	people["John"] = 30.5
	people["Marry"] = 22
	fmt.Printf("%v\n", people) // => map[John:30.5 Marry:22]

	// ✅ Using make()
	map1 := make(map[int]int)
	map1[4] = 8
	fmt.Printf("map1: %#v\n", map1) // -> map1: map[int]int{4:8}

	// ✅ Using map literal with initial values
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
		"CHF": 600, // trailing comma is required in multi-line maps
	}
	fmt.Println(balances)

	// Single-line literal (trailing comma optional)
	m := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m

	// ❌ Duplicate keys not allowed in a map literal:
	// n := map[int]int{1: 3, 1: 4} // ERROR

	// ------------------------------------
	// UPDATING MAPS
	// ------------------------------------
	// Assigning to a key updates the value if it exists,
	// otherwise it inserts a new key:value pair.
	balances["USD"] = 500.5
	balances["GBP"] = 800.8
	fmt.Println(balances)

	// ------------------------------------
	// LOOKUP WITH "COMMA OK" IDIOM
	// ------------------------------------
	// Helps distinguish between:
	//   - missing key
	//   - existing key with zero-value
	v, ok := balances["RON"]
	if ok {
		fmt.Println("The RON Balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	// ------------------------------------
	// ITERATING A MAP
	// ------------------------------------
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	// Since Go 1.12, fmt.Printf prints maps with sorted keys
	fmt.Printf("balances: %v\n", balances)

	// ------------------------------------
	// DELETING ENTRIES
	// ------------------------------------
	delete(balances, "USD")

	// ------------------------------------
	// COMPARING MAPS
	// ------------------------------------
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "X"}

	// ❌ Maps cannot be compared using == (except to nil).
	// fmt.Println(a == b) // ERROR

	// ✅ Trick: Convert maps to string and compare (works if keys/values are strings).
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	// ------------------------------------
	// CLONING A MAP
	// ------------------------------------
	friends := map[string]int{"Dan": 40, "Maria": 35}

	// ⚠️ Assigning copies only the reference (NOT a deep copy).
	neighbors := friends
	friends["Dan"] = 30
	fmt.Println(neighbors) // -> map[Dan:30 Maria:35]

	// ✅ Proper cloning
	colleagues := make(map[string]int)
	for k, v := range friends {
		colleagues[k] = v
	}
	// now colleagues and friends are independent maps
}
