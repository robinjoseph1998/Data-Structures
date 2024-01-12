package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	stringNum := strconv.Itoa(num)
	var numArr []int
	for _, ele := range stringNum {
		oele, _ := strconv.Atoi(string(ele))
		if oele != 0 {
			numArr = append(numArr, oele)
		}
	}
	num1 := ""
	num2 := ""
	allEqual := reflect.DeepEqual(numArr[1:], numArr[:len(numArr)-1])
	if allEqual {
		if len(numArr)%2 != 0 {

		}
	}
	fmt.Println("len(numA r r)", len(numArr))
	if len(numArr) == 2 {
		return numArr[0] + numArr[1]
	}

	sort.Ints(numArr)
	for i, val := range numArr {
		if i%2 == 0 {
			num1 += strconv.Itoa(val)
		} else {
			num2 += strconv.Itoa(val)
		}
	}
	val1, _ := strconv.Atoi(num1)
	val2, _ := strconv.Atoi(num2)
	return val1 + val2

}

func main() {
	num := 999999999
	fmt.Println(splitNum(num))
}
