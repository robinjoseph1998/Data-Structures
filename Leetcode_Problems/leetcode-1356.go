package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortByBits(arr []int) []int {
	type Value struct {
		Value     int
		OnesCount int
	}

	Data := []Value{}
	for _, v := range arr {
		once := 0
		BR := strconv.FormatInt(int64(v), 2)
		for _, one := range BR {
			if one == '1' {
				once++
			}
		}
		Data = append(Data, Value{Value: v, OnesCount: once})
	}
	sort.Slice(Data, func(i, j int) bool {
		if Data[i].OnesCount < Data[j].OnesCount {
			return true
		}
		if Data[i].OnesCount == Data[j].OnesCount {
			return Data[i].Value < Data[j].Value
		}
		return false
	})
	res := []int{}
	for _, v := range Data {
		res = append(res, v.Value)
	}
	return res
}
func main() {
	arr := []int{1000, 1000}
	fmt.Println(sortByBits(arr))

}
