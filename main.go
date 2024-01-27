package main

import (
	"fmt"
)

func countSeniors(details []string) int {
	var ages []string
	for _, Each := range details {
		eachAge := ""
		for i := 0; i < len(Each); i++ {
			if Each[i] == 'M' || Each[i] == 'F' {
				val1 := Each[i+1] - '0'
				val2 := Each[i+2] - '0'
				eachAge += string(val1)
				fmt.Println("val1 string", string(val1))
				eachAge += string(val2)
				ages = append(ages, eachAge)
			}
		}

	}
	fmt.Println("AGES", ages)
	return 0
}

func main() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	fmt.Println(countSeniors(details))

}
