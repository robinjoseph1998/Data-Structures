package main

import (
	"fmt"
)

func digitCount(num string) bool {
	if len(num) == 1 {
		return false
	}
	var intNum []int
	for _, val := range num {
		digit := int(val - '0')
		intNum = append(intNum, digit)
	}
	fmt.Println("intNum", intNum)
	var count int
	for i := 0; i < len(intNum); i++ {
		for k := 0; k < len(intNum); k++ {
			if i == intNum[k] {
				fmt.Println(i, intNum[k])
				count++
			}
		}
		fmt.Println("count", count)
		if count != int(num[i]) {
			return false
		}
	}

	return true
}
func main() {
	num := "1210"
	fmt.Println(digitCount(num))
}
