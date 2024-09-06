package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Binarysearch(arr, 9))
}

func Binarysearch(arr []int, item int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + ((high-low)*(item-arr[low]))/(arr[high]-arr[low])
		if arr[mid] == item {
			return mid
		}
		if item < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
