package solution

func StrStr(haystack string, needle string) int {
	// if needle is empty then return 0
	if len(needle) == 0 {
		return 0
	}

	// cannot start checking for the needle beyond the point where there are
	// not enough characters left in the haystack for the entire needle to fit.
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		// if a match was found, return the starting index
		if match {
			return i
		}
	}

	// if no match found, return -1
	return -1
}
