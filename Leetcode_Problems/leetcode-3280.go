package main

import (
	"fmt"
	"strconv"
)

func convertDateToBinary(date string) string {
	intYear, _ := strconv.Atoi(date[0:4])
	intMonth, _ := strconv.Atoi(date[5:7])
	intDay, _ := strconv.Atoi(date[8:])
	return strconv.FormatInt(int64(intYear), 2) + "-" + strconv.FormatInt(int64(intMonth), 2) + "-" + strconv.FormatInt(int64(intDay), 2)
}

func main() {
	date := "2080-02-29"
	fmt.Println(convertDateToBinary(date))
}
