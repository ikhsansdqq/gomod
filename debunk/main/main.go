package main

import (
	"debunk"
	"fmt"
)

func main() {
	var customArray = []int{1, 3, 5, 8, -2, 2}
	fmt.Println("Quick Sort Method: ")
	debunk.QuickSort(customArray, 0, len(customArray)-1)
	debunk.PrintArray(customArray)

	fmt.Println("Insertion Sort Method: ")
	debunk.InsertionSort(customArray)
	debunk.PrintArray(customArray)
}
