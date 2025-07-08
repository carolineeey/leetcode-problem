package solution

func MinimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// make copy of the last row as the initial state of dp array
	dp := make([]int, len(triangle[len(triangle)-1]))
	copy(dp, triangle[len(triangle)-1])

	// bottom-up calculation
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}
