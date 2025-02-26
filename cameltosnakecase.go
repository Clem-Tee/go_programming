package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {

	if s == "" {
		return ""
	}

	result := []rune{}
	for i, char := range s {
		if i > 0 && char >= 'A' || char <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, char+32)
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
