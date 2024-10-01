package main

import (
	"fmt"
	"strconv"
)

func minPartitions(n string) int {
	max := 0
	for _, v := range n {
		strV := string(v)
		intV, _ := strconv.Atoi(strV)
		if intV > max {
			max = intV
		}
	}
	return max
}

func main() {
	n := "27346209830709182346"
	fmt.Println(minPartitions(n))
}
