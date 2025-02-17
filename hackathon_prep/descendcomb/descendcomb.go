package main

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for x := '9'; x >= '0'; x-- {
				for y := '9'; y >= '0'; y-- {
					num1 := int(i-'0')*10 + int(j-'0')
					num2 := int(x-'0')*10 + int(y-'0')
					if num1 <= num2 {
						continue
					}
					printz01(i)
					printz01(j)
					printz01(' ')
					printz01(x)
					printz01(y)
					if !(i == '0' && j == '1' && x == '0' && y == '0') {
						printz01(',')
						printz01(' ')
					}
				}
			}
		}
	}
	printz01('\n')
}

func printz01(r rune) {
	z01.PrintRune(r)
}

func main() {
	DescendComb()
}
