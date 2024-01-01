package main

import (
	"fmt"
	"math"
)

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	LCM := int(math.Abs(float64(divisor1*divisor2)) / float64(gcd(divisor1, divisor2)))
	TotalCount := uniqueCnt1 + uniqueCnt2

	MaxArr1 := uniqueCnt1 + uniqueCnt1/(divisor1-1)
	if uniqueCnt1%(divisor1-1) == 0 {
		MaxArr1--
	}
	MaxArr2 := uniqueCnt2 + uniqueCnt2/(divisor2-1)
	if uniqueCnt2%(divisor2-1) == 0 {
		MaxArr2--
	}

	MaxBoth := TotalCount + TotalCount/(LCM-1)
	if TotalCount%(LCM-1) == 0 {
		MaxBoth--
	}

	return int(math.Max(float64(MaxArr1), math.Max(float64(MaxArr2), float64(MaxBoth))))
}

func main() {

	divisor1 := 2
	divisor2 := 4
	uniqueCnt1 := 8
	uniqueCnt2 := 2

	fmt.Println(minimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2))

}
