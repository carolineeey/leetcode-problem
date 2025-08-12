package solution

func ConvertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber > 0 {
		offset := (columnNumber - 1) % 26
		res = append(res, 'A'+byte(offset))
		columnNumber = (columnNumber - 1) / 26
	}

	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
