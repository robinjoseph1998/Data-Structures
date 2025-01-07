package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	res := []int32{}
	for _, grade := range grades {
		currGrade := grade
		if grade >= 35 {
			grade++
			for grade%5 != 0 {
				grade++
			}
			if grade-currGrade < 3 {
				res = append(res, grade)
			} else {
				res = append(res, currGrade)
			}
		} else {
			res = append(res, grade)
		}
	}
	return res
}

func main() {
	grades := []int32{73, 67, 38, 33}
	fmt.Println(gradingStudents(grades))
}
