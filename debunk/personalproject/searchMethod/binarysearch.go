package main

import "fmt"

func main() {
	var arr = []int{7, -9, 9, 10}

	fmt.Println(binSearch(arr, 10))
}

func binarySearch(arr []int, l, r, x int) int {
	if r >= 1 {
		mid := l + (r-1)/2
		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}

		return binarySearch(arr, mid+1, r, x)
	}

	return -1
	//	GEEKS FOR GEEKS
}

func binSearch(arr []int, search int) (res, searchCount int) {
	mid := len(arr) / 2
	switch {
	case len(arr) == 0:
		res = -1
	case arr[mid] > search:
		res, searchCount = binSearch(arr[:mid], search)
	case arr[mid] < search:
		res, searchCount = binSearch(arr[mid+1:], search)
		if res >= 0 {
			res += mid + 1
		}
	default:
		res += mid + 1
	}
	searchCount++
	return
}
