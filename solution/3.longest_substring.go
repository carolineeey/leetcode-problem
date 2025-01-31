package solution

func LengthOfLongestSubstring(s string) int {
	temp := make(map[byte]int)
	l := 0
	length := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		// check if the character exists in the temp map
		if _, ok := temp[char]; ok {
			// if it does and the index is greater than l than assign value to l
			if temp[char] >= l {
				l = temp[char] + 1
			}
		}

		// else if the character does not exist, store the index of the character, not the length
		length = max(length, i-l+1)
		temp[char] = i
	}

	return length
}
