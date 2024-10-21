package main

import (
	"fmt"
	"strconv"
)

func largestInteger(num int) int {
	strNum := []byte(strconv.Itoa(num))
	for i := 0; i < len(strNum); i++ {
		for j := i + 1; j < len(strNum); j++ {
			if strNum[i]%2 == strNum[j]%2 && strNum[i] < strNum[j] {
				strNum[i], strNum[j] = strNum[j], strNum[i]
			}
		}
	}
	res, _ := strconv.Atoi(string(strNum))
	return res
}

func main() {
	num := 1234
	fmt.Println(largestInteger(num))
}
