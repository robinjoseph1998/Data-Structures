package main

import "fmt"

func kthCharacter(k int) byte {
	word := "a"
	for len(word) < k {
		newWord := ""
		for _, ch := range word {
			if ch == 'z' {
				newWord += "a"
			} else {
				newWord += string(ch + 1)
			}
		}
		word += newWord
	}
	return word[k-1]
}

func main() {
	k := 5
	fmt.Println(kthCharacter(k))
}
