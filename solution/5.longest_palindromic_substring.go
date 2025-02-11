package solution

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		l++
		r--
		if r-l > end-start {
			start = l
			end = r
		}

		// Check for even-length palindromes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		l++
		r--
		if r-l > end-start {
			start = l
			end = r
		}
	}

	return s[start : end+1]
}
