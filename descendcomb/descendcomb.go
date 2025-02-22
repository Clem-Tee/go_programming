package main

import (
	"fmt"
)

func DescendComb() {
	first := true

	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print("%02d %02d", i, j)
			first = false
		}
	}
	fmt.Println()
}

func main() {
	DescendComb()
}
