package main

import (
	"fmt"
	"strconv"
)

func alternateDigitSum(n int) int {
	s := strconv.Itoa(n)
	var Total int
	for index, val := range s {
		num, _ := strconv.Atoi(string(val))
		if index%2 == 0 {
			Total += num
			fmt.Println("-num", -num)
		} else {
			Total += -num
			fmt.Println("num", num)
		}
	}
	return Total
}

func main() {
	n := 25
	fmt.Println(alternateDigitSum(n))

}
