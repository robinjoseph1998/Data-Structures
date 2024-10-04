package main

import (
	"fmt"
	"strconv"
)

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	if (n & k) != 1 {
		return -1
	}
	diff := n ^ k
	b := strconv.FormatInt(int64(diff), 2)

	ans := 0
	for _, v := range b {
		if v == '1' {
			ans++
		}
	}
	return ans
}

func main() {
	n := 14
	k := 13
	fmt.Println(minChanges(n, k))
}
