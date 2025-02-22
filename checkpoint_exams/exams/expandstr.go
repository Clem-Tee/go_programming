// // program
// takes a string and displays it with exactly three spaces
// btw each word
// no spaces nor tabs at neither the beginning nor the end
// the followed by a newline
// word is a sequence of visible characters
// if no. of args is not 1 or if there are no word
// the program displays nothing

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	word := ""
	var words []string

	for _, char := range input {
		if char != ' ' && char != '\t' {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		return
	}

	result := words[0]
	for i := 1; i < len(words); i++ {
		result += "   " + words[i]
	}

	fmt.Println(result)
}
