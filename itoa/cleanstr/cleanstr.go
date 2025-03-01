package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	start, end := 0, len(input)-1

	for start <= end && (input[start] == ' ' || input[start] == '\t') {
		start++
	}

	for end >= start && (input[end] == ' ' || input[end] == '\t') {
		end--
	}

	if start > end {
		fmt.Println()
		return
	}

	result := ""
	space := false

	for i := start; i <= end; i++ {
		if input[i] == ' ' || input[i] == '\t' {
			if !space {
				result += " "
				space = true
			}
		} else {
			result += string(input[i])
			space = false
		}
	}
	fmt.Println(result)
}
