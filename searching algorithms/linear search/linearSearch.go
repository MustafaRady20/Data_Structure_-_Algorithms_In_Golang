package main

import "fmt"

func main() {

	list := []int{1, 2, 3, 4, 5, 8, 61, 8, 7, 6}
	fmt.Println(linearSearch(list, 5))
}

func linearSearch(list []int, value any) (any, int) {
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return list[i], i
		}
	}
	return nil, -1
}
