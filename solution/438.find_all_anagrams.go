package solution

func FindAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	// the length of s cannot be less than p
	if sLen < pLen {
		return []int{}
	}

	// initialize frequency map
	var sCount, pCount [26]int
	for i := 0; i < pLen; i++ {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}

	// helper function to compare the frequency
	matches := func() bool {
		for i := 0; i < 26; i++ {
			if sCount[i] != pCount[i] {
				return false
			}
		}
		return true
	}

	// slide the window over string s
	var result []int
	for i := 0; i <= sLen-pLen; i++ {
		if i > 0 {
			// Add new character to the window
			sCount[s[i+pLen-1]-'a']++
			// Remove the character left out of the window
			sCount[s[i-1]-'a']--
		}

		if matches() {
			result = append(result, i)
		}
	}

	return result
}
