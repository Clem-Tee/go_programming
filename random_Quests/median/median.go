package main

import (
	"fmt"
	"sort"
)

// findMedian calculates the median of an integer slice
func findMedian(arr []int) float64 {
	// Step 1: Sort the array using Go's built-in sort function
	sort.Ints(arr)

	// Step 2: Find the length of the array
	n := len(arr)

	// Step 3: Find the median
	// If the number of elements is odd, return the middle element
	if n%2 == 1 {
		return float64(arr[n/2])
	} else {
		// If the number of elements is even, return the average of the two middle elements
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
}

func main() {
	numbers := []int{7, 3, 1, 5, 9, 9}
	median := findMedian(numbers)
	fmt.Println("Median:", median)
}
