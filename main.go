package main

import (
	"fmt"
)

func maximumNumberOfStringPairs(words []string) int {
	Count := 0
	j := 0
	for i := 0; i < len(words)-1; i++ {
		rev := ""
		val := string(words[i])

		for i := len(val) - 1; i >= 0; i-- {
			rev += string(val[i])
		}

		for j = i + 1; j < len(words); j++ {
			if rev == string(words[j]) {
				fmt.Println("::::::::::::::REV", rev)
				fmt.Println("(words[j])", string(words[j]))
				Count++
				break
			}
		}
		if j == len(words)-1 {
			fmt.Println("j", j)
			break
		}
	}
	return Count
}

func main() {

	words := []string{"cd", "ac", "dc", "ca", "zz"}

	fmt.Println(maximumNumberOfStringPairs(words))

}
