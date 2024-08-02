package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if CheckPalindrome(s) {
		return s
	}
	max := 0
	bigP := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if CheckPalindrome(s[i:j]) {
				if len(s[i:j]) > max {
					max = len(s[i:j])
					bigP = s[i:j]
				}
			}
		}
	}
	return bigP
}

func CheckPalindrome(text string) bool {
	size := len(text)
	for i := 0; i < size/2; i++ {
		if text[i] != text[size-1-i] {

			return false
		}
	}
	return true

}

func main() {
	s := "abb"
	fmt.Println(longestPalindrome(s))
}
