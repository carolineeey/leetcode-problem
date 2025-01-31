package solution

func LengthOfLongestSubstring(s string) int {
	temp := make(map[byte]int)
	l := 0
	length := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if _, ok := temp[char]; ok {
			if temp[char] >= l {
				l = temp[char] + 1
			}
		}
		length = max(length, i-l+1)
		temp[char] = i
	}

	return length
}
