package main

import "fmt"

func countCharacters(words []string, chars string) int {
	wordMap := make(map[rune]int)
	for _, v := range chars {
		wordMap[v]++
	}
	totalLength := 0
	for _, word := range words {
		if canFormword(word, wordMap) {
			totalLength += len(word)
		}
	}
	return totalLength
}

func canFormword(word string, wordMap map[rune]int) bool {
	wordCount := make(map[rune]int)
	for _, c := range word {
		wordCount[c]++
		if wordCount[c] > wordMap[c] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))
}
