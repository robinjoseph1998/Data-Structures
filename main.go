package main

import (
	"fmt"
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	strNum := ""
	for i := 0; i < len(num); i++ {
		strNum += strconv.Itoa(num[i])
	}
	fmt.Println("strNum", strNum)
	intNum, _ := strconv.ParseUint(strNum, 10, 64)
	fmt.Println("intNum", intNum)
	sum := intNum + uint64(k)
	fmt.Println("sum", sum)
	strSum := strconv.FormatUint(sum, 10)
	var arr []int
	for _, val := range strSum {
		intv, _ := strconv.Atoi(string(val))
		arr = append(arr, intv)
	}
	return arr
}

func main() {

	num := []int{3, 8, 0, 3, 0, 2, 7, 0, 7, 6, 4, 9, 9, 1, 7, 6, 6, 1, 6, 4}
	k := 670
	fmt.Println(addToArrayForm(num, k))

}
