package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	strNum := []rune(str)
	for i := 0; i < len(strNum); i++ {
		if strNum[i] == '6' {
			strNum[i] = '9'
			break
		}
	}
	res := string(strNum)
	intRes, _ := strconv.Atoi(res)
	return intRes
}

func main() {
	num := 9669
	fmt.Println(maximum69Number(num))
}
