package main

import (
	"fmt"
	"strconv"
)

func isBalanced(num string) bool {

	oddSum, evenSum := 0, 0
	for idx, v := range num {
		if idx%2 == 0 {
			val, _ := strconv.Atoi(string(v))
			oddSum += val
		} else {
			val, _ := strconv.Atoi(string(v))
			evenSum += val
		}
	}

	return oddSum == evenSum
}

func main() {
	num := "1234"
	fmt.Println(isBalanced(num))
}
