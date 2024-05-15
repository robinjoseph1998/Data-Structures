package main

import (
	"fmt"
	"unicode"
)

func numberOfSpecialChars(word string) int {
	count := 0
	smallest := make(map[rune]bool)
	capital := make(map[rune]bool)
	for _, char := range word {
		if char >= 'A' && char <= 'Z' {
			capital[char] = true
		} else {
			smallest[char] = true
		}
	}

	for char := range capital {
		if smallest[unicode.ToLower(char)] {
			count++
		}
	}
	return count
}

func main() {
	word := "ABCabc"
	fmt.Println(numberOfSpecialChars(word))
}
