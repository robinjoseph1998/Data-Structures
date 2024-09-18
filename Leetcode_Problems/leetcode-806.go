package main

import (
	"fmt"
)

func numberOfLines(widths []int, s string) []int {
	words := make(map[rune]int)
	j := 0
	for i := 'a'; i <= 'z'; i++ {
		words[i] = widths[j]
		j++
	}
	count := 0
	lines := 1
	for _, v := range s {
		width := words[v]
		if count+width > 100 {
			lines++
			count = width
		} else {
			count += width
		}
	}

	return []int{lines, count}
}

func main() {
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"
	fmt.Println(numberOfLines(widths, s))
}
