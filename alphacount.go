package main

import "fmt"

func AlphaCount(s string) int {
	count := 0
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'A' && runeS[i] <= 'Z' {
			count++
		}
		if runeS[i] >= 'a' && runeS[i] <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(AlphaCount("hello world"))
	fmt.Println(AlphaCount("I wanna go home"))
	fmt.Println(AlphaCount("mimiKemmy 123hls"))

}
