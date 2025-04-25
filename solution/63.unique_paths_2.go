package solution

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// If starting point is an obstacle
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	grid[0][0] = 1

	// fill first column
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && grid[i-1][0] == 1 {
			grid[i][0] = 1
		}
	}

	// fill first row
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && grid[0][j-1] == 1 {
			grid[0][j] = 1
		}
	}

	// fill the rest
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	// return the number of ways to reach the cell
	return grid[m-1][n-1]
}
