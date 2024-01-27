package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	count := 0
	for _, EachVal := range details {
		age, _ := strconv.Atoi(EachVal[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}

func main() {
	details := []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"}
	fmt.Println(countSeniors(details))

}
