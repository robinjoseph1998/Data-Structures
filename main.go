package main

import (
	"fmt"
)

func sortString(s string) string {
	var temp []int
	for i := 0; i < len(s); i++ {
		temp = append(temp, int(s[i]))
	}
	for i := 0; i < len(s); i++ {
		if temp[i] > temp[i+1] {

		}

	}
	result := ""
	for i := 0; i < len(s); i++ {
		result += string(temp[i])
	}
	return result
}

func main() {
	s := "rat"
	fmt.Println(sortString(s))
}
