package main

import (
	"fmt"
)

func HashCode(dec string) string {
	size := len(dec)
	hashed := ""

	for _, char := range dec {
		hashVal := (int(char) + size) % 127
		if hashVal < 32 {
			hashVal += 33
		}
		hashed += string(rune(hashVal))
	}
	return hashed
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
