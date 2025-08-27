package data

// BinarySearch performs a binary search on a sorted slice of integers.
// 🔹 needle: the value to search for
// 🔹 haystack: a sorted slice of integers
// Returns true if needle exists in the slice, false otherwise.
//
// ⚠️ Important: The slice must be sorted in ascending order before calling this function.
func BinarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	// Binary search loop
	for low <= high {
		// Calculate the middle index
		median := (low + high) / 2

		if haystack[median] < needle {
			// If middle value is less than needle, search in the right half
			low = median + 1
		} else {
			// Otherwise, search in the left half
			high = median - 1
		}
	}

	// Check if low points to the needle
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
