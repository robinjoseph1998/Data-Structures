package main

import (
	"fmt"
	"strings"
)

func hasPrefixandSuffix(a, b string) bool {
	return strings.HasPrefix(b, a) && strings.HasSuffix(b, a)
}

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if hasPrefixandSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func main() {
	words := []string{"a", "aba", "ababa", "aa"}
	fmt.Println(countPrefixSuffixPairs(words))
}
