package solution

func UniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// iterate over the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// if we are at the first row or column that means only one way to reach the cell
			if i == 0 || j == 0 {
				grid[i][j] = 1
			} else {
				// memoize the number of ways
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	// return the number of ways to reach the cell
	return grid[m-1][n-1]
}
