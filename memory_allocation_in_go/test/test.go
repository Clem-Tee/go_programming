package main

import "fmt"

func createArrayOfSize(size int) []int {

	// var answer []int = make([]int, size)

	var answer []int

	for i := 0; i < size; i++ {

		answer = append(answer, i+1)
	}
	return answer
}

func main() {

	size := 5
	myarray := createArrayOfSize(size)

	fmt.Println(myarray)
}
