package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	if len(num1) > 9 || len(num2) > 9 {
		return MultBigInt(num1, num2)
	}

	return strconv.FormatInt(int64(n1)*int64(n2), 10)
}

func MultBigInt(num1, num2 string) string {
	bigNum1, _ := new(big.Int).SetString(num1, 10)
	bigNum2, _ := new(big.Int).SetString(num2, 10)

	return new(big.Int).Mul(bigNum1, bigNum2).String()
}

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2))
}
