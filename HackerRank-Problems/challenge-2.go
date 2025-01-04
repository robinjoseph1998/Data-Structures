package main

import "fmt"

func plusMinus(arr []int32) {
	positiveCount, negativeCount, zeroCount := 0, 0, 0

	for _, num := range arr {
		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("%.6f\n", float32(positiveCount)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(negativeCount)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(zeroCount)/float32(len(arr)))

}

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}
