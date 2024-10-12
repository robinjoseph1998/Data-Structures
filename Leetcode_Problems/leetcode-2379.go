package main

import (
	"fmt"
	"math"
)

func minimumRecolors(blocks string, k int) int {
	blocksArr := []rune(blocks)
	min := math.MaxInt
	for i := 0; i <= len(blocksArr)-k; i++ {
		currK := CountWhitblock(blocksArr[i : i+k])
		if currK < min {
			min = currK
		}
	}
	return min
}

func CountWhitblock(arr []rune) int {
	count := 0
	for _, v := range arr {
		if v == 'W' {
			count++
		}
	}
	return count
}

func main() {
	blocks := "WBBWWBBWBW"
	k := 7
	fmt.Println(minimumRecolors(blocks, k))
}
