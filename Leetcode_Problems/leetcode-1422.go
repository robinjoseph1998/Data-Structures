package main

import "fmt"

func maxScore(s string) int {
	totalOnes := 0
	for _, v := range s {
		if v == '1' {
			totalOnes++
		}
	}
	maxScore := 0
	leftZeros := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			leftZeros++
		} else {
			totalOnes--
		}
		currentScore := leftZeros + totalOnes
		if currentScore > maxScore {
			maxScore = currentScore
		}
	}
	return maxScore
}

func main() {
	s := "011101"
	fmt.Println(maxScore(s))
}
