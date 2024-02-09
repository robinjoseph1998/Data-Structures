package main

import (
	"fmt"
	"strconv"
)

func lastVisitedIntegers(words []string) []int {
	count := 0
	for _, each := range words {
		if each == "prev" {
			count++
		}
	}
	var result []int
	for i := len(words) - 1; i > 0; i-- {
		if words[i] != "prev" {
			val, _ := strconv.Atoi(words[i])
			result = append(result, val)
			count--
		}
		// if i < count {
		// 	result = append(result, -1)
		// }
	}
	return result
}

func main() {
	words := []string{"1", "2", "prev", "prev", "prev"}
	fmt.Println(lastVisitedIntegers(words))
}
