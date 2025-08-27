package data

// LinearSearch searches for a key in a slice using a simple linear scan.
//
// 🔹 datalist: slice of integers to search
// 🔹 key: the value to search for
//
// Returns true if the key is found, false otherwise.
//
// Time Complexity: O(n) in worst case (searches every element)
// Space Complexity: O(1) (in-place)
//
// ✅ Best for small or unsorted datasets, or when elements are not uniformly distributed.
func LinearSearch(datalist []int, key int) bool {
	// Iterate through each element
	for _, item := range datalist {
		if item == key {
			// Return immediately if key is found
			return true
		}
	}
	// Key not found
	return false
}
