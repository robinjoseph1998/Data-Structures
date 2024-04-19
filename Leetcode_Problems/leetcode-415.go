package main

import (
	"fmt"
	"math/big"
)

func addStrings(num1 string, num2 string) string {
	int1, _ := new(big.Int).SetString(num1, 10)
	int2, _ := new(big.Int).SetString(num2, 10)
	return new(big.Int).Add(int1, int2).String()
}

func main() {
	num1 := "3876620623801494171"
	num2 := "6529364523802684779"
	fmt.Println(addStrings(num1, num2))
}
