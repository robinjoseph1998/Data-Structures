package main

import (
	"fmt"
)

func isSubstringPresent(s string) bool {
	sub := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		sub[s[i:i+2]] = true
	}
	reverseStr := reverse(s)
	for i := 0; i < len(reverseStr)-1; i++ {
		if sub[reverseStr[i:i+2]] {
			return true
		}
	}
	return false
}
func reverse(str string) string {
	runes := []rune(str)
	j := 0
	for i := len(str) - 1; i >= len(str)/2; i-- {
		runes[i], runes[j] = runes[j], runes[i]
		j++
	}
	return string(runes)
}

func main() {
	s := "leetcode"
	fmt.Println(isSubstringPresent(s))
}
