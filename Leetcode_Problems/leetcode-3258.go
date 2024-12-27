package main

import "fmt"

func countKConstraintSubstrings(s string, k int) int {
	count := 0
	for i := 0; i < len(s); i++ {
		zero_count, one_count := 0, 0
		for j := i; j < len(s); j++ {
			if s[j] == '1' {
				one_count++
			} else {
				zero_count++
			}
			if zero_count > k && one_count > k {
				break
			}
			count++
		}
	}
	return count
}
func main() {

	//Example: 1
	s := "10101"
	k := 1
	fmt.Println(countKConstraintSubstrings(s, k))

	//Example: 2
	s = "1010101"
	k = 2
	fmt.Println(countKConstraintSubstrings(s, k))

	//Example: 3
	s = "11111"
	k = 1
	fmt.Println(countKConstraintSubstrings(s, k))

}
