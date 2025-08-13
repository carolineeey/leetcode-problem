package solution

func TitleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		val := int(columnTitle[i]-'A') + 1
		result = result*26 + val
	}

	return result
}
