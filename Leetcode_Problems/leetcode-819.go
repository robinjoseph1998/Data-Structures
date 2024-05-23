package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	if len(banned) == 0 {
		re := regexp.MustCompile(`[^\w]+`)
		words := re.Split(paragraph, -1)
		word := strings.ToLower(words[0])
		return word
	}
	Lp := strings.ToLower(paragraph)
	re := regexp.MustCompile(`[^\w]+`)
	words := re.Split(Lp, -1)
	Occurences := make(map[string]int)
	is := 0
	for _, each := range words {
		for i := 0; i < len(banned); i++ {
			if each == banned[i] {
				is++
			}
		}
		if is == 0 {
			Occurences[each]++
		}
		is = 0
	}
	count := 0
	res := ""
	for key, val := range Occurences {
		if count < val {
			count = val
			res = key
		}

	}
	return res
}

func main() {
	paragraph := "Bob. hIt, baLl"
	banned := []string{"bob", "hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}
