package solution

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	// 1D dynamic programming to store results for each row
	dp := make([]int, cols+1) // +1 to the cols to make sure j-1 or i-1 doesn't go out of bounds
	maxLen := 0
	prev := 0 // prev stores the value of the top-left (diagonal) neighbor before it was overwritten

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			// dp[j]: length of the largest square ending at matrix[i-1][j-1]
			temp := dp[j] // store the current cell before updating it, so we can pass it to prev for the next column
			if matrix[i-1][j-1] == '1' {
				// dp[j] : top  	--> square ending directly above
				// dp[j-1] : left  	--> square ending to the left
				// prev : top-left 	--> square ending diagonally above-left
				dp[j] = findMin(dp[j], dp[j-1], prev) + 1 // update dp[j] with min of neighbors
				if dp[j] > maxLen {
					maxLen = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return maxLen * maxLen
}

func findMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	}

	if b < c {
		return b
	}
	return c
}
