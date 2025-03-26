package solution

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	//7, 6, 4, 3, 1
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i] //update min price if current value is less than minPrice as it will make minus profit
		} else {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		}
	}

	return maxProfit
}
