package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isFascinating(n int) bool {
	a := n * 2
	fmt.Println("a", a)
	b := n * 3
	fmt.Println("b", b)
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)
	strN := strconv.Itoa(n)

	concatVal := strN + strA + strB
	result := true
	var conArr []int
	for _, each := range concatVal {
		conArr = append(conArr, int(each-'0'))
	}
	sort.Ints(conArr)
	count := 0
	for i := 0; i < len(conArr); i++ {
		if i+1 == conArr[i] {
			fmt.Println("i+1", i+1)
			fmt.Println("concat Value i", conArr[i])
			count++
		} else {
			result = false
		}
	}

	return result
}

func main() {
	n := 192
	fmt.Println(isFascinating(n))

}
