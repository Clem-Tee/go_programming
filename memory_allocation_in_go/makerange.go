package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slice := make([]int, max-min)
	for i := min; i < max; i++ {
		slice[i-min] = i
	}
	return slice
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
