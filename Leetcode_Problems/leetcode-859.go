package main

import (
	"fmt"
)

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		charCount := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			charCount[s[i]]++
			if charCount[s[i]] > 1 {
				return true
			}
		}
		return false
	}

	var diff []int

	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}

	if len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]] {
		return true
	}
	return false
}

func main() {
	s := "aaaaaaabc"
	goal := "aaaaaaacb"
	fmt.Println(buddyStrings(s, goal))
}
