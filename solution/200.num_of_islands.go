package solution

func NumIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])

	if row == 0 && col == 0 {
		return 0
	}

	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= row || c < 0 || c >= col || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0' // mark as visited so we don't count it again
		// recursively visit all its neighboring cells (up, down, left, right) if they are also '1'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				count++
				// continue doing this until there are no more connected '1's — meaning we’ve explored one full island.
				dfs(r, c)
			}
		}
	}

	return count
}
