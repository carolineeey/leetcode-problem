package solution

func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowZero := make([]bool, m)
	colZero := make([]bool, n)

	// record columns and rows to zero
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	// set matrix[i][j] to zero if its row or column is marked
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}
