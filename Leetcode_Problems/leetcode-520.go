package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func detectCapitalUse(word string) bool {
	if isAllCaps(word) {
		return true
	}
	if isAllSmall(word) {
		return true
	}
	firstUpper := isUpper(word[0])
	secondUpper := isUpper(word[1])

	if !firstUpper && secondUpper {
		return false
	}
	res := false
	if firstUpper {
		for idx, v := range word {
			if idx != 0 {
				if unicode.IsLower(v) {
					res = true
				} else {
					res = false
					break
				}
			}
		}
	}
	if !res {
		for idx, v := range word {
			if idx != 0 {
				if unicode.IsUpper(v) {
					res = true
				} else {
					res = false
					break
				}
			}
		}
	}
	return res
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
func isAllCaps(s string) bool {
	var allCapsPattern = "^[A-Z]+$"
	re := regexp.MustCompile(allCapsPattern)
	return re.MatchString(s)
}
func isAllSmall(s string) bool {
	var allCapsPattern = "^[a-z]+$"
	re := regexp.MustCompile(allCapsPattern)
	return re.MatchString(s)
}

func main() {
	str := "FlaG"
	fmt.Println(detectCapitalUse(str))
}
