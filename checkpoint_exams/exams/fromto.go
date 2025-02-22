package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	// Check for invalid range
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	// Determine step direction
	step := 1
	if from > to {
		step = -1
	}

	// Generate the range string
	for i := from; i != to; i += step {
		result += fmt.Sprintf("%02d, ", i)
	}
	result += fmt.Sprintf("%02d\n", to)

	return result
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
