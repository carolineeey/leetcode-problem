package solution

func GeneratePascalTriangle(rowIndex int) []int {
	wholeRes := make([][]int, rowIndex)

	for i := 0; i < rowIndex; i++ {
		wholeRes[i] = make([]int, i+1)        // create row with i+1 elements
		wholeRes[i][0], wholeRes[i][i] = 1, 1 // first and last elements are always 1

		for j := 1; j < i; j++ {
			wholeRes[i][j] = wholeRes[i-1][j-1] + wholeRes[i-1][j]
		}
	}

	return wholeRes[rowIndex-1]
}

func GeneratePascalTriangle2(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}
