package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	runeS := []rune(s)
	result := ""

	for _, char := range runeS {
		if char >= 'a' && char <= 'z' {
			for i := 0; i < int(char-'a'+1); i++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			for i := 0; i < int(char-'A'+1); i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
