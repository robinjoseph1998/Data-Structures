package main

import "fmt"

func longestPalindrome(s string) int {
	n := len(s)
	pairs := 0
	count := map[rune]int{}
	for _, ch := range s {
		count[ch]++
		if count[ch]%2 == 0 {
			pairs++
		}
	}
	if pairs*2 == n {
		return n
	}
	return pairs*2 + 1
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
