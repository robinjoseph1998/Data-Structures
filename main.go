package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestGoodInteger(num string) string {
	var numInt []int
	for i := 0; i < len(num); i++ {
		val, _ := strconv.Atoi(string(num[i]))
		numInt = append(numInt, val)
	}
	var Triples []int
	for i := 0; i < len(numInt); i++ {
		EachstrVal := strconv.Itoa(numInt[i])
		TriplledStr := ""
		for i := 1; i <= 3; i++ {
			TriplledStr += EachstrVal
		}
		if strings.Contains(num, TriplledStr) {
			tripleIntegers, _ := strconv.Atoi(TriplledStr)
			Triples = append(Triples, tripleIntegers)
		}
	}
	if len(Triples) == 0 {
		return ""
	}
	sort.Ints(Triples)
	val := Triples[len(Triples)-1]
	if val == 0 {
		return "000"
	}
	str := strconv.Itoa(val)
	return str
}

func main() {

	num := "2300019"

	fmt.Println(largestGoodInteger(num))

}
