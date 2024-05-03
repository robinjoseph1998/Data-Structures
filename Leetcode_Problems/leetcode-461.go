package main

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	bitX := strconv.FormatInt(int64(x), 2)
	bitY := strconv.FormatInt(int64(y), 2)

	maxLen := len(bitX)
	if len(bitY) > maxLen {
		maxLen = len(bitY)
	}
	bitX = fmt.Sprintf("%0*b", maxLen, x)
	bitY = fmt.Sprintf("%0*b", maxLen, y)
	diff := 0
	for i := 0; i < maxLen; i++ {
		if bitX[i] != bitY[i] {
			diff++
		}
	}
	return diff
}

func main() {
	x := 680142203
	y := 1111953568
	fmt.Println(hammingDistance(x, y))

}
