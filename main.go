package main

import (
	"fmt"
)

func findTheLongestBalancedSubstring(s string) int {
	var Zeroes, Once, TotalCount int = 0, 0, 0
	Flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			Zeroes++
			Once = 0
		}
		if s[i] == '1' {
			Once++
			Flag = true
		}
		if Zeroes >= Once && TotalCount < Once {
			TotalCount = Once
		}
		if Flag && s[i] == '0' {
			Zeroes = 1
			Flag = false
		}
	}
	return TotalCount * 2
}

func main() {

	s := "010"

	fmt.Println(findTheLongestBalancedSubstring(s))

}
