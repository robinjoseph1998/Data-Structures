package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for idx, w := range words {
		reversedWord := reverse(w)
		words[idx] = reversedWord
	}
	return strings.Join(words, " ")
}

func reverse(word string) string {
	words := []byte(word)
	j := len(words) - 1
	for i := 0; i < len(word)/2; i++ {
		words[i], words[j] = words[j], words[i]
		j--
	}
	return string(words)
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
