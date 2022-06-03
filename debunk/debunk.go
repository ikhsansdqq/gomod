package debunk

import (
	"fmt"
	"strconv"
)

func PrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}

	fmt.Println("")
}

func Partition(arr []int, low, high int) int {
	var (
		pivot = arr[low]
		i     = low
		j     = high
	)

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var (
			temp         = arr[i]
			holePosition = i
		)

		for holePosition > 0 && arr[holePosition-1] > temp {
			arr[holePosition] = arr[holePosition-1]
			holePosition--
		}

		arr[holePosition] = temp
	}
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		var pivot = Partition(arr, low, high)
		QuickSort(arr, low, pivot)
		QuickSort(arr, pivot+1, high)
	}
}

// COPYRIGHT NOTICE
// ALL THE PACKAGE SOURCES ARE BELONGS TO IKHSAN ASSIDIQIE WHO CREATE THE PACKAGE
// FREE TO BE USED FOR PERSONAL PROJECT OR COMMERCIAL USE BUT DO NOT CHANGE THE NAME OF THE CREATOR
