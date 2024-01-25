package main

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) int {
	seperator := func(num int) (bool, int) {
		fmt.Println("num", num)
		if num == 10 {
			return true, 100
		}
		sqr := num * num
		if num == 36 {
			fmt.Println("SQR 36 ", sqr)
		}
		strSqr := strconv.Itoa(sqr)
		total := 0
		for _, val := range strSqr {
			intVal := val - '0'
			total += int(intVal)
		}
		if num == 36 {
			fmt.Println("Total", total)
		}
		if total == num {
			return true, sqr
		}
		return false, 0
	}
	var PunishmentNumber int
	for i := 1; i <= n; i++ {
		if isSepartor, value := seperator(i); isSepartor {
			fmt.Println(":::::::::::::::i", i)
			fmt.Println("value", value)
			PunishmentNumber += value
		}
	}
	return PunishmentNumber
}

func main() {
	n := 37
	fmt.Println(punishmentNumber(n))
}
