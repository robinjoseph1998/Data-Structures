package main

import (
	"fmt"
	"sort"
)

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	var Store1 []int
	var Store2 []int
	var Runner int
	for i := 1; i <= 1000; i++ {
		if i%divisor1 != 0 {
			Store1 = append(Store1, i)
			Runner++
		}
		if Runner == uniqueCnt1 {
			break
		}

	}
	fmt.Println("Store1", Store1)
	Runner = 0
	for i := 1; i <= 1000; i++ {
		if i%divisor2 != 0 {

			Store2 = append(Store2, i)
			Runner++
		}
		if Runner == uniqueCnt2 {
			break
		}
	}
	fmt.Println("Store2", Store2)
	Contains := func(Store1 []int, value int) bool {
		for _, val := range Store1 {
			if val == value {
				return true
			}
		}
		return false
	}

	for _, value := range Store2 {
		if Contains(Store1, value) {
			for Contains(Store1, value) || value%divisor2 != 0 {
				value++
			}
		}
		Store1 = append(Store1, value)
	}
	sort.Ints(Store1)
	fmt.Println("Store", Store1)
	return Store1[len(Store1)-1]
}

func main() {

	divisor1 := 5
	divisor2 := 5
	uniqueCnt1 := 9
	uniqueCnt2 := 3

	fmt.Println(minimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2))

}
