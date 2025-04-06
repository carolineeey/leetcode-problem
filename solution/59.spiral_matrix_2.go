package solution

func GenerateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	// initialize an n x n matrix filled with zeroes
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	rowBegin, rowEnd := 0, n-1
	columnBegin, columnEnd := 0, n-1
	num := 1

	for rowBegin <= rowEnd && columnBegin <= columnEnd {
		// traverse right
		for i := columnBegin; i <= columnEnd; i++ {
			matrix[rowBegin][i] = num
			num++
		}
		rowBegin++

		// traverse down
		for i := rowBegin; i <= rowEnd; i++ {
			matrix[i][columnEnd] = num
			num++
		}
		columnEnd--

		// traverse left
		if rowBegin <= rowEnd {
			for i := columnEnd; i >= columnBegin; i-- {
				matrix[rowEnd][i] = num
				num++
			}
			rowEnd--
		}

		// traverse up
		if columnBegin <= columnEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				matrix[i][columnBegin] = num
				num++
			}
			columnBegin++
		}
	}

	return matrix
}
