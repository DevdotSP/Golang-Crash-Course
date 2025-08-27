package data

// Insertionsort sorts a slice of integers in ascending order using the Insertion Sort algorithm.
//
// 🔹 items: slice of integers to be sorted
//
// Insertion Sort builds the final sorted array one item at a time.
// It compares the current element with previous elements and inserts it in the correct position.
//
// Time Complexity: O(n^2) worst case, O(n) best case (already sorted)
// Space Complexity: O(1) (in-place sorting)
// ✅ Efficient for small datasets or nearly sorted arrays. Not recommended for large unsorted datasets.
func Insertionsort(items []int) {
	n := len(items)

	// Iterate over the array starting from the second element
	for i := 1; i < n; i++ {
		j := i

		// Move the current element backward until it is in the correct position
		for j > 0 && items[j-1] > items[j] {
			// Swap elements if they are out of order
			items[j-1], items[j] = items[j], items[j-1]
			j = j - 1
		}
	}
}
