package solution

func Solve(board [][]byte) {
	row, col := len(board), len(board[0])

	if row == 0 || col == 0 {
		return
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r == row || c == col || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// capture unsurrounded regions (O -> T)
	for r := 0; r < row; r++ {
		dfs(r, 0)
		dfs(r, col-1)
	}

	for c := 0; c < col; c++ {
		dfs(0, c)
		dfs(row-1, c)
	}

	// flip remaining O to X and T back to O
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'T' {
				board[r][c] = 'O'
			}
		}
	}
}
