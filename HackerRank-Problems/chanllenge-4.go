package main

import (
	"fmt"
)

func birthdayCakeCandles(candles []int32) int32 {
	cmap := make(map[int32]int)
	for _, each := range candles {
		cmap[each]++
	}
	max := 0
	count := 0
	for k, v := range cmap {
		if k > int32(max) {
			max = int(k)
			count = v
		}
	}
	return int32(count)
}

func main() {
	arr := []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}

