package main

import (
	"fmt"
)

// Abort function calculates the average of any numbers of integers

// func Abort(a, b, c, d, e int) int {
// 	return (a + b + c + d + e) / 5
// }

func Abort(num ...int) float64 {

	if float64(len(num)) == 0 {
		return 0 // avoid division by zero
	}

	sum := 0

	for _, nums := range num {
		sum += nums
	}
	return float64(sum) / float64(len(num))
}

func main() {
	middle := Abort(0)
	fmt.Println(middle)
}
