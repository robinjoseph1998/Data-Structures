package main

import (
	"fmt"
	"strconv"
)

func minBitFlips(start int, goal int) int {
	Flip := 0
	startBinary := strconv.FormatInt(int64(start), 2)
	goalBinary := strconv.FormatInt(int64(goal), 2)
	if len(goalBinary) < len(startBinary) {
		for len(goalBinary) < len(startBinary) {
			tmp := "0"
			tmp += goalBinary
			goalBinary = tmp
		}
	} else if len(goalBinary) > len(startBinary) {
		for len(goalBinary) > len(startBinary) {
			tmp := "0"
			tmp += startBinary
			startBinary = tmp
		}
	}
	for i := 0; i < len(startBinary); i++ {
		if startBinary[i] != goalBinary[i] {
			Flip++
		}
	}
	return Flip
}

func main() {
	start := 99
	goal := 29
	fmt.Println(minBitFlips(start, goal))
}
