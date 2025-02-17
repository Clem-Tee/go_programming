package main

import (
	"fmt"
)

func LoafOfBread(str string) string {
	s := []rune(str)

	if len(s) < 5 {
		return "Invalid Output\n"
	}

	var r []rune
	count := 0

	for _, char := range s {
		if char != ' ' {
			r = append(r, char)
			count++
		}
		if count%5 == 0 && count != 0 {
			r = append(r, ' ')
		}
	}
	return string(r) + "\n"
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
