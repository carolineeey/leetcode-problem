package solution

func MaxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// If k is large enough, we can consider it as infinite transactions
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	// DP Approach: profit[i][j] represents max profit with at most i transactions at day j
	profit := make([][]int, k+1)
	for i := range profit {
		profit[i] = make([]int, n)
	}

	for t := 1; t <= k; t++ {
		maxDiff := -prices[0] // Represents the max difference found so far
		for d := 1; d < n; d++ {
			profit[t][d] = max(profit[t][d-1], prices[d]+maxDiff)
			maxDiff = max(maxDiff, profit[t-1][d]-prices[d])
		}
	}

	return profit[k][n-1]
}
