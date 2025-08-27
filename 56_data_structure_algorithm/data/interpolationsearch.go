package data

// InterpolationSearch searches for a key in a sorted slice using the Interpolation Search algorithm.
//
// 🔹 array: a sorted slice of integers (ascending order required)
// 🔹 key: the value to search for
//
// Returns the index of the first occurrence of the key if found.
// If the key is not found, it returns the index where the key would be inserted.
//
// Interpolation Search improves upon binary search by estimating the position
// based on the value of the key relative to the first and last elements.
// Best used when elements are uniformly distributed.
//
// Time Complexity:
// - Best case: O(log log n) (uniform distribution)
// - Worst case: O(n) (non-uniform distribution)
// Space Complexity: O(1) (in-place)
func InterpolationSearch(array []int, key int) int {
	min, max := array[0], array[len(array)-1]
	low, high := 0, len(array)-1

	for {
		// Key is smaller than minimum element, insert at beginning
		if key < min {
			return low
		}

		// Key is larger than maximum element, insert at end
		if key > max {
			return high + 1
		}

		// Estimate the position of the key (interpolation formula)
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// Check if the guessed position contains the key
		if array[guess] == key {
			// Scan backward to find the first occurrence in case of duplicates
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		// Adjust search range
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}
