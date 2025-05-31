package solution

import "unicode"

func IsCharacterPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// move left pointer to next alphanumeric character
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		// move right pointer to previous alphanumeric character
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		// compare characters (case-insensitive)
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
