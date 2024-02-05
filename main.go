package main

import (
	"fmt"
)

func maximumOddBinaryNumber(s string) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		if count > 1 || i == len(s)-1 {
			result += "1"
			count--
		} else if count == 1 {
			result += "0"
		}
	}
	return result
}

func main() {
	s := "010"
	fmt.Println(maximumOddBinaryNumber(s))
}
