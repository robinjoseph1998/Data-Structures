package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	var l, w int
	sqrt := int(math.Sqrt(float64(area)))

	for i := sqrt; i > 0; i-- {
		if area%i == 0 {
			l = area / i
			w = i
			break
		}
	}
	return []int{l, w}
}

func main() {
	area := 4
	fmt.Println(constructRectangle(area))
}
