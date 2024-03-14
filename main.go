package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	sum := 0
	strNumb := strconv.Itoa(n)
	for i := 0; i < len(strNumb); i++ {
		val, _ := strconv.Atoi(string(strNumb[i]))
		sum += val * val
	}
	if sum == 1 {
		return true
	} else if sum == 4 {
		return false
	} else {
		return isHappy(sum)
	}
}

func main() {

	fmt.Println(isHappy(19))

}
