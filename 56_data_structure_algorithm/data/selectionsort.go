package data

// Selectionsort sorts a slice of integers in ascending order using the Selection Sort algorithm.
//
// 🔹 items: slice of integers to be sorted
//
// Selection Sort works by repeatedly finding the minimum element from the unsorted portion
// and swapping it with the first unsorted element.
//
// Time Complexity: O(n^2) for all cases (best, average, worst)
// Space Complexity: O(1) (in-place sorting)
// ✅ Simple and intuitive, but not efficient for large datasets.
func Selectionsort(items []int) {
	n := len(items)

	// Iterate over each position in the slice
	for i := 0; i < n; i++ {
		minIdx := i

		// Find the index of the minimum element in the unsorted portion
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum element with the first unsorted element
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
