package data

// Combsort sorts a slice of integers in ascending order using the Comb Sort algorithm.
//
// 🔹 items: slice of integers to be sorted
//
// Comb Sort is an improvement over Bubble Sort.
// It eliminates turtles, or small values near the end of the list, by using a gap
// to compare and swap elements that are far apart, then shrinking the gap each iteration.
//
// Time Complexity: O(n^2) worst case, O(n log n) average case
// Space Complexity: O(1) (in-place sorting)
// ✅ More efficient than Bubble Sort, but less common than quicksort or mergesort.
func Combsort(items []int) {
	n := len(items)
	gap := n          // Initial gap size is the length of the slice
	shrink := 1.3     // Shrink factor determines how the gap decreases
	swapped := true   // Flag to track if any swaps occurred in a pass

	// Continue iterations until no swaps occur
	for swapped {
		swapped = false

		// Shrink the gap for the next iteration
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}

		// Compare elements with the current gap
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				// Swap elements if they are in wrong order
				items[i], items[i+gap] = items[i+gap], items[i]
				swapped = true
			}
		}
	}
}
