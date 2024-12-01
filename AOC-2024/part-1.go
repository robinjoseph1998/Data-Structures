package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findDiff(inp1, inp2 []int) int {
	sort.Ints(inp1)
	sort.Ints(inp2)
	length := 0 // or length of inp2
	sum := 0
	for i := 0; i < length; i++ {
		sum += abs(inp1[i], inp2[i])
	}
	return sum
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	input := `Add Input here`

	lines := strings.TrimSpace(input)
	lineArr := strings.Split(lines, "\n")
	res := 0

	left := []int{}
	right := []int{}
	for _, eachLine := range lineArr {
		numbers := strings.Fields(eachLine)
		if len(numbers) == 2 {

			leftID, _ := strconv.Atoi(numbers[0])
			left = append(left, leftID)

			rightID, _ := strconv.Atoi(numbers[1])
			right = append(right, rightID)
		}
	}
	res += findDiff(left, right)
	fmt.Println("res", res)
}
