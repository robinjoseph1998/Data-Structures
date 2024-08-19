package main

import (
	"fmt"
	"unicode"
)

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	letterCount := 0
	isCapital := false
	isZerotoNine := false
	isVowel := false
	notVowel := false
	for _, v := range word {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			return false
		} else if unicode.IsLetter(v) {
			letterCount++
			if unicode.IsUpper(v) {
				isCapital = true
			}
			if vowels[string(v)] {
				isVowel = true
			} else if unicode.IsLetter(v) && !vowels[string(v)] {
				notVowel = true
			}
		} else if unicode.IsDigit(v) && v < 10 || v >= 0 {
			isZerotoNine = true
		}
	}
	if notVowel && isVowel {
		return true
	}
	return isCapital && isZerotoNine && isVowel && notVowel && letterCount >= 3
}

func main() {
	word := "aya"
	fmt.Println(isValid(word))
}
