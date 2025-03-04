package solution

func IsValidSudoku(board [][]byte) bool {
	// use map to track occurrences of digits in rows, columns, and subgrids
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	subgrids := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		subgrids[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			// calculate subgrid index
			subgridIdx := (i/3)*3 + j/3

			// check for duplicates in row, column, and subgrid
			if rows[i][num] || columns[j][num] || subgrids[subgridIdx][num] {
				return false
			}

			// mark the digit as seen
			rows[i][num] = true
			columns[j][num] = true
			subgrids[subgridIdx][num] = true
		}
	}

	return true
}
