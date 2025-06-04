package solution

func NumTrees(n int) int {
	// Dynamic Programming is a technique to solve problems by breaking them into smaller
	// subproblems, storing their results, and reusing those results to avoid redundant computation.

	// The number of unique BSTs with n nodes follows the Catalan number (dynamic programming):
	// G(n)=G(0)*G(n−1)+G(1)*G(n−2)+...+G(n−1)*G(0)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
