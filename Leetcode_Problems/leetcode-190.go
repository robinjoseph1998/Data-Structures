package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	BinaryVal := fmt.Sprintf("%032b", num)
	runes := []rune(BinaryVal)
	j := 0
	for i := len(runes) - 1; i >= len(runes)/2; i-- {
		runes[i], runes[j] = runes[j], runes[i]
		j++
	}
	reversedBinary := string(runes)
	ans, _ := strconv.ParseUint(reversedBinary, 2, 32)
	return uint32(ans)
}

func main() {
	var n uint32 = 43261596
	fmt.Println(reverseBits(n))
}
