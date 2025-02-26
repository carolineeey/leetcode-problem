package solution

// only add open parentheses if open < n
// only add close parentheses if close < open
// valid if close == open == n
func GenerateParenthesis(n int) []string {
	var result []string
	var backtrack func(string, int, int)

	backtrack = func(current string, open, close int) {
		// if the current string has n open and n close parentheses, it's valid
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		// add open parentheses if we haven't use all open parentheses
		if open < n {
			backtrack(current+"(", open+1, close)
		}

		// add close parentheses if we can match it open parentheses
		if close < open {
			backtrack(current+")", open, close+1)
		}

	}

	// start recursion with empty string and 0 value of open and close parentheses
	backtrack("", 0, 0)

	return result
}
