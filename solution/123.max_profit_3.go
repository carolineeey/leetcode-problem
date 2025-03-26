package solution

func MaxProfit3(prices []int) int {
	n := len(prices)
	if n < 0 {
		return 0
	}

	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := -prices[0]
	secondSell := 0

	for i := 1; i < n; i++ {
		firstBuy = max(firstBuy, -prices[i])              // Best time to buy the first stock
		firstSell = max(firstSell, prices[i]+firstBuy)    // Best time to sell after first buy
		secondBuy = max(secondBuy, firstSell-prices[i])   // Best time to buy the second stock
		secondSell = max(secondSell, prices[i]+secondBuy) // Best time to sell after second buy
	}

	return secondSell
}
