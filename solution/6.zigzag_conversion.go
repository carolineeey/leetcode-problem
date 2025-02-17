package solution

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	i := 0
	d := 1 // the direction, 1 is for going down, -1 is for going up
	rows := make([]string, numRows)

	for _, char := range s {
		rows[i] += string(char)
		if i == 0 {
			d = 1 // move down
		} else if i == numRows-1 {
			d = -1 // move up
		}
		i += d
	}

	// concatenate all rows to make the final string
	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
