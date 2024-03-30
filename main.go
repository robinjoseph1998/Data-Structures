package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	fmt.Println("Channels in golang")

// 	myCh := make(chan int, 2)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func(ch <-chan int, wg *sync.WaitGroup) {
// 		for val := range ch {
// 			fmt.Println("val", val)
// 		}
// 		wg.Done()
// 	}(myCh, wg)

// 	go func(ch chan<- int, wg *sync.WaitGroup) {
// 		myCh <- 8
// 		myCh <- 9
// 		close(ch)
// 		wg.Done()
// 	}(myCh, wg)

// 	wg.Wait()
// }

import (
	"fmt"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {
	strVal := ""
	itsFloat := false
	lowestBoundary := -2147483648
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' || s[i] == '.' {
			strVal += string(s[i])
			if s[i] == '.' {
				itsFloat = true
			}
		} else if unicode.IsLetter(rune(s[i])) && strVal == "" {
			return 0
		}
	}
	var intVal int
	if itsFloat {
		intVal, _ := strconv.ParseFloat(strVal, 64)

		s := strconv.FormatFloat(intVal, 'f', -1, 64)
		f := s[0]
		res, _ := strconv.Atoi(string(f))
		return res

	} else {
		intVal, _ = strconv.Atoi(strVal)
	}
	if intVal < lowestBoundary {
		return lowestBoundary
	}
	return intVal
}

func main() {
	s := "3.14159"
	fmt.Println(myAtoi(s))
}
