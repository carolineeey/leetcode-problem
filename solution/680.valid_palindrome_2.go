package solution

func ValidPalindrome(s string) bool {
	isPalind := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// try skipping left or right character
			return isPalind(left+1, right) || isPalind(left, right-1)
		}

		left++
		right--
	}

	return true
}
