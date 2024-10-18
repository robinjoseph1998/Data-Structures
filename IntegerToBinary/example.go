package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42                                     // Example integer
	binaryStr := strconv.FormatInt(int64(num), 2) // Convert to binary string
	fmt.Println(binaryStr)                        // Output: "101010"
}
