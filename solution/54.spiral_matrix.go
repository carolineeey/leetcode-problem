package solution

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rowBegin, rowEnd := 0, len(matrix)-1
	columnBegin, columnEnd := 0, len(matrix[0])-1
	var spiralOrder []int

	for rowBegin <= rowEnd && columnBegin <= columnEnd {
		// Traverse right
		for i := columnBegin; i <= columnEnd; i++ {
			spiralOrder = append(spiralOrder, matrix[rowBegin][i])
		}
		rowBegin++

		// Traverse down
		for i := rowBegin; i <= rowEnd; i++ {
			spiralOrder = append(spiralOrder, matrix[i][columnEnd])
		}
		columnEnd--

		// Traverse left (ensure we are still within valid row bounds)
		if rowBegin <= rowEnd {
			for i := columnEnd; i >= columnBegin; i-- {
				spiralOrder = append(spiralOrder, matrix[rowEnd][i])
			}
			rowEnd--
		}

		// Traverse up (ensure we are still within valid column bounds)
		if columnBegin <= columnEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				spiralOrder = append(spiralOrder, matrix[i][columnBegin])
			}
			columnBegin++
		}
	}

	return spiralOrder
}
