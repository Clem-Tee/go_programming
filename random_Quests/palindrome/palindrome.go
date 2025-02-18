package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", "")) // Normalize by removing spaces and converting to lowercase
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	testString := "No lemon"
	fmt.Println("Is palindrome:", isPalindrome(testString))
}
