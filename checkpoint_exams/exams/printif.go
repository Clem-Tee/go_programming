package main

import (
	"fmt"
)

func PrintIf(str string) string {
	len := len(str)

	if len >= 3 || len == 0 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
