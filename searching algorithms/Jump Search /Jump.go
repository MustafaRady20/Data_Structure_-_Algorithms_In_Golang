package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(jumpSearch(arr, 8))
}

func jumpSearch(arr []int, item int) int {
	start := 0
	blockSize := int(math.Sqrt(float64(len(arr))))
	end := blockSize

	// find the right block to search though it
	for arr[end] <= item && end < len(arr) {
		start = end
		end = end + blockSize
		if end > len(arr)-1 {
			end = len(arr)
		}
	}

	// search for the element in the block
	for i := start; i < end; i++ {
		if arr[i] == item {
			return i
		}
	}
	return -1
}
