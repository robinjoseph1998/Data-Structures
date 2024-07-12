package main

import (
	"fmt"
	"strconv"
)

func bitwiseComplement(n int) int {
	binarN := fmt.Sprintf("%b", n)
	compliment := ""
	for _, v := range binarN {
		if v == '0' {
			compliment += "1"
		} else {
			compliment += "0"
		}
	}
	res, _ := strconv.ParseInt(compliment, 2, 64)
	return int(res)
}

func main() {
	n := 5
	fmt.Println(bitwiseComplement(n))
}
