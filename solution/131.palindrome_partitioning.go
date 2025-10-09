package solution

func Partition(s string) [][]string {
	var res [][]string
	var path []string

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			// reach the end, valid partition found
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s[start : end+1]) {
				path = append(path, s[start:end+1])
				backtrack(end + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
