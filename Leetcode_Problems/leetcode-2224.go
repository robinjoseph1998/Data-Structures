package main

import (
	"fmt"
	"time"
)

func convertTime(current string, correct string) int {
	currentTime, _ := time.Parse("15:04", current)
	correctTime, _ := time.Parse("15:04", correct)

	diff := int(correctTime.Sub(currentTime).Minutes())
	if diff < 0 {
		diff += 1440
	}
	minutes := []int{60, 15, 5, 1}
	operations := 0
	for _, v := range minutes {
		operations += diff / v
		diff %= v
	}
	return operations
}

func main() {
	current := "02:30"
	correct := "04:35"
	fmt.Println(convertTime(current, correct))
}
