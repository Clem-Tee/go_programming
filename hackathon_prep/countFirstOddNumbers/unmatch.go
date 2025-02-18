package main

import (
	"fmt"
)

func Unmatch(a []int) int {
	countMap := make(map[int]int)

	// Count occurrences of each number
	for _, num := range a {
		countMap[num]++
	}

	// Find the first number with an odd count
	for _, num := range a {
		if countMap[num]%2 != 0 {
			return num
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 2, 2, 3, 3, 3, 1, 2, 4, 3, 4, 4, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
