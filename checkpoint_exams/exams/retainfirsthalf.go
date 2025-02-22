package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	size := len(str)

	if size == 0 {
		return ""
	}
	if size == 1 {
		return str
	}
	halfSize := size / 2
	return str[:halfSize]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
