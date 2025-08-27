package data

// Bubblesort sorts a slice of integers in ascending order using the bubble sort algorithm.
//
// 🔹 items: slice of integers to be sorted
//
// Bubble sort repeatedly steps through the list, compares adjacent elements,
// and swaps them if they are in the wrong order.
// The process repeats until no swaps are needed, indicating the slice is sorted.
//
// Time Complexity: O(n^2) on average and worst case
// Space Complexity: O(1) (in-place sorting)
// ✅ Best for small datasets or educational purposes; not efficient for large datasets.
func Bubblesort(items []int) {
	n := len(items)
	sorted := false

	// Continue iterating until the slice is fully sorted
	for !sorted {
		swapped := false // Tracks if any swaps occurred in this pass

		// Compare adjacent elements
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				// Swap if elements are in wrong order
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}

		// If no swaps occurred, the slice is sorted
		if !swapped {
			sorted = true
		}

		// Reduce the range since the last element is already in correct position
		n = n - 1
	}
}
