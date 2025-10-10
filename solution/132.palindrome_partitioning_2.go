package solution

func MinCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	// dp[i] = min cuts needed for substring s[0:i+1]
	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}

	// palindrome table
	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || isPal[j+1][i-1]) {
				isPal[j][i] = true
				if j == 0 {
					dp[i] = 0 // entire substring s[0:i+1] is palindrome
				} else {
					if dp[j-1]+1 < dp[i] {
						dp[i] = dp[j-1] + 1
					}
				}
			}
		}
	}

	return dp[n-1]
}
