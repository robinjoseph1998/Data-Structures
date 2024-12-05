package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("./day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the grid from the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	n := len(lines)    // Number of rows
	m := len(lines[0]) // Number of columns

	// Check if "XMAS" condition is satisfied at position (i, j)
	hasXMAS := func(i, j int) bool {
		// Boundary check to avoid out-of-bounds
		if !(1 <= i && i < n-1 && 1 <= j && j < m-1) {
			return false
		}

		// The center character must be "A"
		if lines[i][j] != 'A' {
			return false
		}

		// Check both diagonals
		diag1 := string([]byte{lines[i-1][j-1], lines[i+1][j+1]}) // Top-left to bottom-right
		diag2 := string([]byte{lines[i-1][j+1], lines[i+1][j-1]}) // Top-right to bottom-left

		return (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM")
	}

	// Count occurrences
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if hasXMAS(i, j) {
				ans++
			}
		}
	}

	// Print the result
	fmt.Println(ans)
}
