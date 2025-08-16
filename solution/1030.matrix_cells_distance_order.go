package solution

import "sort"

func AllCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	var cells [][]int

	// collect all cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cells = append(cells, []int{r, c})
		}
	}

	// sort by Manhattan distance
	sort.Slice(cells, func(i, j int) bool {
		d1 := abs(cells[i][0]-rCenter) + abs(cells[i][0]-cCenter)
		d2 := abs(cells[j][0]-rCenter) + abs(cells[j][1]-cCenter)
		return d1 < d2
	})

	return cells
}
