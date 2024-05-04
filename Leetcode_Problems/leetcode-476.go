package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	br := fmt.Sprintf("%b", num)
	rev := ""
	for i := 0; i < len(br); i++ {
		if br[i] == '1' {
			rev += "0"
		} else {
			rev += "1"
		}
	}
	decimal, _ := strconv.ParseInt(rev, 2, 64)
	return int(decimal)

}

func main() {
	num := 5
	fmt.Println(findComplement(num))

}
