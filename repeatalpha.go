package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	var result string

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			var repeatcount int
			if char >= 'a' && char <= 'z' {
				repeatcount = int(char - 'a' + 1)
			} else {
				repeatcount = int(char - 'A' + 1)
			}
			for i := 0; i < repeatcount; i++ {
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
