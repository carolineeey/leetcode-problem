package solution

func CheckInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	// the length of s1 cannot be greater than s2
	if s1Len > s2Len {
		return false
	}

	// init frequency map
	var s1Count, s2Count [26]int
	for i := 0; i < s1Len; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	// helper function to compare
	matches := func() bool {
		for i := 0; i < 26; i++ {
			if s1Count[i] != s2Count[i] {
				return false
			}
		}
		return true
	}

	// sliding windows
	for i := 0; i <= s2Len-s1Len; i++ {
		if i > 0 {
			s2Count[s2[i-1]-'a']--
			s2Count[s2[i+s1Len-1]-'a']++
		}
		if matches() {
			return true
		}
	}
	return false
}
