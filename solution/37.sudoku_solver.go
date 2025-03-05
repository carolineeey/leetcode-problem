package solution

func SolveSudoku(board [][]byte) {
	solveSudoku(board)
}

func solveSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValidSudoku(board, i, j, num) {
						board[i][j] = num
						if solveSudoku(board) {
							return true
						}
						board[i][j] = '.' // backtracking
					}
				}
				return false
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte, row, col int, num byte) bool {
	boxRowStart := (row / 3) * 3
	boxColStart := (col / 3) * 3

	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num || board[boxRowStart+i/3][boxColStart+i%3] == num {
			return false
		}
	}

	return true
}
