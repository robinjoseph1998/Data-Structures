package main

import (
	"fmt"
)

func isSubstringPresent(s string) bool {
	if isPalindrome(s) {
		return true
	}
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] != s[i+1] {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func isPalindrome(text string) bool {
	size := len(text)
	for i := 0; i < size/2; i++ {
		if text[i] != text[size-1-i] {

			return false
		}
	}
	return true
}

func main() {
	s := "abcba"
	fmt.Println(isSubstringPresent(s))

}
