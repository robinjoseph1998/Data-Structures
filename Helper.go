package main

// import (
// 	"fmt"
// 	"sort"
// )

// func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
// 	var Store1 []int
// 	var Store2 []int
// 	var Runner int
// 	for i := 1; i <= 1000; i++ {
// 		if i%divisor1 != 0 {
// 			Store1 = append(Store1, i)
// 			Runner++
// 		}
// 		if Runner == uniqueCnt1 {
// 			break
// 		}

// 	}
// 	fmt.Println("Store1", Store1)
// 	Runner = 0
// 	for i := 1; i <= 1000; i++ {
// 		if i%divisor2 != 0 {

// 			Store2 = append(Store2, i)
// 			Runner++
// 		}
// 		if Runner == uniqueCnt2 {
// 			break
// 		}
// 	}
// 	Contains := func(arr []int, value int) bool {
// 		for _, val := range arr {
// 			if val == value {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	var Divisor int
// 	if divisor1%2 != 0 {
// 		Divisor = divisor1
// 	} else {
// 		Divisor = divisor2
// 	}
// 	for _, value := range Store2 {
// 		for Contains(Store1, value) {
// 			value++
// 			for value%Divisor == 0 {
// 				fmt.Println("value", value)
// 				value++
// 			}
// 		}
// 		Store1 = append(Store1, value)
// 	}
// 	sort.Ints(Store1)

// 	return Store1[len(Store1)-1]
// }

// func main() {

// 	divisor1 := 2
// 	divisor2 := 4
// 	uniqueCnt1 := 8
// 	uniqueCnt2 := 2

// 	fmt.Println(minimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2))

// }
