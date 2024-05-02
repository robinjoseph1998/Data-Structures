package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	wordToPattern := make(map[string]byte)
	patternToWord := make(map[byte]string)

	for i, word := range words {
		p := pattern[i]

		if prevWord, exists := patternToWord[p]; exists {
			if prevWord != word {
				return false
			}
		} else {
			patternToWord[p] = word
		}

		if prevPattern, exists := wordToPattern[word]; exists {
			if prevPattern != p {
				return false
			}
		} else {
			wordToPattern[word] = p
		}
	}
	return true
}
func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}
