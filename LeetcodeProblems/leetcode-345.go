package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	var idxVovels []int
	vovels := "aeiouAEIOU"
	for idx, char := range s {
		if strings.ContainsRune(vovels, char) {
			idxVovels = append(idxVovels, idx)
		}
	}
	runes := []rune(s)

	for i := 0; i < len(idxVovels)/2; i++ {
		j := len(idxVovels) - i - 1

		runes[idxVovels[i]], runes[idxVovels[j]] = runes[idxVovels[j]], runes[idxVovels[i]]
	}
	return string(runes)
}

func main() {

	s := "leetcode"
	fmt.Println(reverseVowels(s))

}
