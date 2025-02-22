package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	count := 0

	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	s := "World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
