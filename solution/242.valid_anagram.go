package solution

func IsAnagramV1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sortedString := sortString(t)
	sortedAnagram := sortString(s)

	if sortedAnagram != sortedString {
		return false
	}

	return true
}

func IsAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// map to count the frequency of each character
	charCount := make(map[rune]int)

	// increment the count for each character in s
	for _, char := range s {
		charCount[char]++
	}

	// decrement the count for each character in t
	for _, char := range t {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	// if all counts are zero than s and t are valid anagram
	return true
}
