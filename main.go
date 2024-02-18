package main

import "fmt"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {

	s := "ababa"
	fmt.Println(removePalindromeSub(s))
}
