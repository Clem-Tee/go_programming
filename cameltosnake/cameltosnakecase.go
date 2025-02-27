package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var result []rune
	for i, ch := range s {
		if i > 0 && ch >= 'A' && ch <= 'Z' {
			result = append(result, '_', ch)
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
