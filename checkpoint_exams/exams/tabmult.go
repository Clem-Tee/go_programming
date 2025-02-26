package main

import (
	"fmt"
	"os"
)

func toInt(s string) int {
	if s == "" {
		return -1
	}
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	n := toInt(os.Args[1])
	if n <= 0 {
		fmt.Println()
		return
	}

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, i*n)
	}
}
