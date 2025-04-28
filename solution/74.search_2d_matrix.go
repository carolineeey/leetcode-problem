package solution

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

func SearchMatrixWithBinary(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	start, end := 0, rows*cols-1
	for start <= end {
		mid := (start + end) / 2
		row, col := mid/cols, mid%cols
		current := matrix[row][col]

		if current == target {

			return true
		} else if current > target {
			end = mid - 1
		} else if current < target {
			start = mid + 1
		}
	}

	return false
}
