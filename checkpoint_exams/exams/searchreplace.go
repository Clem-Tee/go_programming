package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	inputString := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	if len(oldChar) != 1 || len(newChar) != 1 {
		return
	}

	modifiedString := ""
	for _, char := range inputString {
		if string(char) == oldChar {
			modifiedString += newChar
		} else {
			modifiedString += string(char)
		}
	}
	fmt.Println(modifiedString)
}
