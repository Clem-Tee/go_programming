package main

import (
	"fmt" // Used to print text to the screen
	"os"  // Used to get input from the command line
)

// Function to change uppercase letters to lowercase and vice versa
func switchCase(s string) string {
	result := []rune(s) // Convert the input string into a list of characters

	for i := 0; i < len(result); i++ { // Loop through each character
		if result[i] >= 'a' && result[i] <= 'z' { // If the character is a lowercase letter
			result[i] = result[i] - 32 // Convert it to uppercase by subtracting 32
		} else if result[i] >= 'A' && result[i] <= 'Z' { // If the character is an uppercase letter
			result[i] = result[i] + 32 // Convert it to lowercase by adding 32
		}
	}
	return string(result) // Convert the modified list back to a string and return it
}

func main() {
	if len(os.Args) != 2 { // Check if exactly one argument is provided
		return // If not, do nothing
	}
	fmt.Println(switchCase(os.Args[1])) // Call switchCase and print the result
}
