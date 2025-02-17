package main

import (
	"fmt"
	"strings"
)

func FoodDeliveryTime(order string) int {

	// Convert input to lowercase to make it case-insensitive
	order = strings.ToLower(order)

	prepTimes := map[string]int{
		"burger":  35,
		"chips":   10,
		"nuggets": 12,
	}

	if time, exists := prepTimes[order]; exists {
		return time
	}
	return -1
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("BURger"))
	fmt.Println(FoodDeliveryTime("pizza"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("Chips") + FoodDeliveryTime("nuGGets"))
}
