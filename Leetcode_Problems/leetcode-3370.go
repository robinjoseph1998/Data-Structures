package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallestNumber(n int) int {
	for {
		binary := strconv.FormatInt(int64(n), 2)
		if !strings.Contains(binary, "0") {
			return n
		}
		n++
	}
}

func main() {
	n := 5
	fmt.Println(smallestNumber(n))
}
