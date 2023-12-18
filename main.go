package main

import (
	"fmt"
	"math"
)

func repeatedCharacter(s string) byte {
	// for i := 0; i < len(s); i++ {
	// 	if i+1 < len(s) {
	// 		if s[i] == s[i+1] {
	// 			return s[i]
	// 		}
	// 	}
	// }
	var val int = math.MaxInt64
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j < val {
					val = j
				}
			}

		}

	}
	fmt.Println(string(s[val]))
	return byte(s[val])
}

func main() {

	s := "izspyzhxvhmvsqekrauyugcbepvifvgnpthxrqunslwvgfdnzfzdxockaoomqybnsfzewkcspwpepvbyohccnoivagjhzplnkcvr"

	fmt.Println(repeatedCharacter(s))

}
