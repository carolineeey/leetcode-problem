package solution

func CountNegatives(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])

	total := 0
	current := column

	for i := 0; i < row; i++ {
		for current-1 >= 0 && grid[i][current-1] < 0 {
			current--
		}
		total += column - current
	}

	return total
}
