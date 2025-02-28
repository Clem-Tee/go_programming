package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	checkNext := true // Start by checking the first letter
	for _, char := range s {
		if checkNext && char >= 'A' && char <= 'Z' {
			return true // Found a capital letter at the start of a word
		}
		checkNext = char == ' ' || char == '!' || char == '?' || char == '.' // Reset after spaces or punctuation
	}

	return false
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?")) // false
	fmt.Println(IsCapitalized("Hello How Are You"))   // true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))   // true
	fmt.Println(IsCapitalized("Whatsthis4"))          // true
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))      // true
	fmt.Println(IsCapitalized(""))                    // false
}
