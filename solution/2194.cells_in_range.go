package solution

func CellsInRange(s string) []string {
	startCol := s[0]
	startRow := int(s[1] - '0')
	endCol := s[3]
	endRow := int(s[4] - '0')

	var res []string
	for c := startCol; c <= endCol; c++ {
		for r := startRow; r <= endRow; r++ {
			cell := string(c) + string(rune(r+'0'))
			res = append(res, cell)
		}
	}

	return res
}
