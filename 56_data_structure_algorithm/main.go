package main

import (
	"fmt"
	"master_go_programming/56_data_structure_algorithm/data"
)

func main() {

	//Linearsearch
	// items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	// fmt.Println(data.Linearsearch(items, 78))

	//Binarysearch
	// items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	// fmt.Println(data.BinarySearch(63, items))

	//Interpolationsearch
	// items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	// fmt.Println(data.InterpolationSearch(items, 63))

	//Bubblesort
	// slice := helper.GenerateSlice(20)
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	// data.Bubblesort(slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	//Quicksort
	// slice := helper.GenerateSlice(20)
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	// data.Quicksort(slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	//selectionsort
	// slice := helper.GenerateSlice(20)
	// fmt.Println("This is selectionsort")
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	// data.Selectionsort(slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	//combsort
	// slice := helper.GenerateSlice(20)
	// fmt.Println("this is combsort")
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	// data.Combsort(slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	//insertionsort
	// slice := helper.GenerateSlice(20)
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	// data.Insertionsort(slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	list := &data.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	fmt.Println("Initial List: ")
	data.PrintList(list)

	list.Remove(2)
	fmt.Println("List afteR Removing 2: ")
	data.PrintList(list)

	list.Remove(4)
	fmt.Println("List afteR Removing 4: ")
	data.PrintList(list)
}
