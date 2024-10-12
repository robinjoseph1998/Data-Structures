package main

import (
	"fmt"
)

func secondsToRemoveOccurrences(s string) int {
	time := 0
	runeS := []rune(s)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(runeS); i++ {
			if runeS[i] == '1' && runeS[i-1] == '0' {
				runeS[i], runeS[i-1] = runeS[i-1], runeS[i]
				swapped = true
				i++
			}
		}
		if swapped {
			time++
		}
	}

	ss := string(runeS)
	fmt.Println("ss", ss)
	return time
}

func main() {
	s := "0110101"
	fmt.Println(secondsToRemoveOccurrences(s))
}
