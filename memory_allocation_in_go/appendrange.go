package main

import (
	"fmt"
)

func AppendRange(min, max int) []int {
	slice := []int{}

	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		slice = append(slice, i)
	}
	return slice
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
