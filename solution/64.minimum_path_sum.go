package solution

func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// fill first column
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// fill first row
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// fill the rest
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	// return the number of ways to reach the cell
	return grid[m-1][n-1]
}
