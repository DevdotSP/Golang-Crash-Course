package data

import (
	"math/rand"
	"time"
)

// Quicksort sorts a slice of integers in ascending order using the Quick Sort algorithm.
//
// 🔹 a: slice of integers to be sorted
//
// Quick Sort is a divide-and-conquer sorting algorithm:
// 1. Select a pivot element (randomly for better performance on nearly sorted data).
// 2. Partition the array into elements less than pivot and greater than pivot.
// 3. Recursively sort the subarrays.
//
// Time Complexity:
// - Average case: O(n log n)
// - Worst case: O(n^2) (rare with random pivot)
// Space Complexity: O(log n) due to recursion
//
// ✅ Efficient for large datasets and widely used in practice.
func Quicksort(a []int) []int {
	if len(a) < 2 {
		// Base case: array with 0 or 1 element is already sorted
		return a
	}

	// Set a random seed
	rand.Seed(time.Now().UnixNano())

	// Choose a random pivot index
	pivot := rand.Intn(len(a))

	// Swap the pivot with the rightmost element
	left, right := 0, len(a)-1
	a[pivot], a[right] = a[right], a[pivot]

	// Partition the array around the pivot
	for i := range a[:right] {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	// Swap the pivot into its correct position
	a[left], a[right] = a[right], a[left]

	// Recursively sort the subarrays
	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}
