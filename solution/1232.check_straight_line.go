package solution

func CheckStraightLine(coordinates [][]int) bool {
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (y1-y0)*(x-x0) != (x1-x0)*(y-y0) {
			return false
		}
	}
	return true
}

// y1-y0 / y-y0 = x1-x0 / x-x0
// y1-y0 * x-x0 = x1-x0 * y-y0
