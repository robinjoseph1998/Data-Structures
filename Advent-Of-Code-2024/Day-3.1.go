package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func FindValidMul(str string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := regex.FindAllStringSubmatch(str, -1)
	sum := 0
	for _, v := range matches {
		val1, _ := strconv.Atoi(v[1])
		val2, _ := strconv.Atoi(v[2])

		sum += val1 * val2
	}

	return sum
}

func main() {
	str := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	fmt.Println(FindValidMul(str))

}
