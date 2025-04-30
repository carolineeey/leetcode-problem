package solution

func SearchMatrix2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows == 0 {
		return false
	}

	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}

	return false
}
