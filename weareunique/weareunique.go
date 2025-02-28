package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	charMap := make(map[rune]int)

	for _, char := range str1 {
		charMap[char] = 1
	}

	for _, char := range str2 {
		if _, exists := charMap[char]; exists {
			charMap[char] = 2
		} else {
			charMap[char] = 1
		}
	}
	count := 0
	for _, value := range charMap {
		if value == 1 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
