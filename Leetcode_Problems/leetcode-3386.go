package main

import (
	"fmt"
	"math"
)

func buttonWithLongestTime(events [][]int) int {

	max := math.MinInt64
	prevTime := 0
	key := 0

	for _, event := range events {

		CurrTime := event[1]

		if abs(CurrTime, prevTime) > max {

			max = abs(CurrTime, prevTime)
			key = event[0]

		} else if abs(CurrTime, prevTime) == max {

			if event[0] < key {

				max = abs(CurrTime, prevTime)
				key = event[0]
			}
		}
		prevTime = CurrTime
	}
	return key
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	events := [][]int{{1, 5}, {19, 9}, {6, 10}, {6, 11}, {16, 14}, {1, 16}, {15, 19}}
	fmt.Println(buttonWithLongestTime(events))
}
