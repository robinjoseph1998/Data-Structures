package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	var compressed []byte
	count := 1
	for i := 0; i < len(chars); i++ {
		if i+1 < len(chars) && chars[i] == chars[i+1] {
			count++
		} else {
			compressed = append(compressed, chars[i])
			if count > 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					compressed = append(compressed, countStr[j])
				}
			}
			count = 1
		}
	}
	if len(compressed) <= len(chars) {
		copy(chars, compressed)
		return len(compressed)
	}
	return len(chars)
}

func main() {

	chars := []byte{61, 62, 63}

	fmt.Println(compress(chars))

}
