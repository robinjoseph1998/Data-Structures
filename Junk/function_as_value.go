package main

import (
	"fmt"
	"math"
)

func Compute(fn func(float64, float64) float64) float64 {
	return fn(100, 60)
}

func main() {

	Engine := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(Engine(300, 200))

	fmt.Println(Compute(Engine))
	fmt.Println(Compute(math.Pow))

}
