package main

// import "fmt"

// func detailer(val ...interface{}) {

// 	for _, values := range val {

// 		switch val := values.(type) {

// 		case string:
// 			fmt.Println("Name", val)
// 		case int:
// 			fmt.Println("Age", val)
// 		case []int:
// 			var total int
// 			for _, marks := range val {
// 				total += marks
// 			}
// 			avg := total/len(val) - 1
// 			fmt.Println("Avg Marks", avg)
// 		}
// 	}
// }

// func main() {

// 	datas := map[string]interface{}{

// 		"Name":  "Robin Joseph",
// 		"Age":   25,
// 		"Marks": []int{100, 200, 300, 100, 400},
// 	}

// 	detailer(datas["Name"], datas["Age"], datas["Marks"])

// }
