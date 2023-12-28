package main

import (
	"fmt"
)

func captureForts(forts []int) int {
	fmt.Println("LENGTH", len(forts))
	var Count, First, Last, Max, Runner int
	for k := Runner; k < len(forts); k++ {
		if forts[k] == 0 && k-1 >= 0 {
			if forts[k-1] != 0 {
				First = forts[k-1]
				for forts[k] == 0 && k < len(forts)-1 {
					Count++
					k++
				}
			}
			if forts[k] != 0 {
				Last = forts[k]
			}
			if First != Last {
				fmt.Println(":::First:", First, "and  Last:", Last)
				if Count > Max {
					Max = Count
					Count = 0
					Runner = k
					fmt.Println("------------------------:::::MAX:::::", k, "::::::::", Max)
				}
			} else {
				Runner = k
				Count = 0
			}
		}
		Count = 0
	}
	return Max
}

func main() {

	forts := []int{-1, 0, 1, 0, 0, 0}

	fmt.Println(captureForts(forts))
}
