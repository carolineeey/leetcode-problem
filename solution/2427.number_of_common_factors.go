package solution

func CommonFactors(a int, b int) int {
	m := a
	if a < b {
		m = b
	}
	count := 0
	for i := 1; i <= m; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}
