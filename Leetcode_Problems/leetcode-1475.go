func finalPrices(prices []int) []int {
	res := []int{}
	for i := 0; i < len(prices); i++ {
		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}
		res = append(res, prices[i]-discount)
	}
	return res
}


func main() {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}
