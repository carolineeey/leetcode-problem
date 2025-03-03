package solution

func LongestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	// initiate dummy value to the stack with a base index of -1
	stack := []int{-1}
	maxValue := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// push the index of the opening parentheses onto the stack
			stack = append(stack, i)
		} else {
			// pop the top element for the closing parentheses
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				// calculate the length of the valid substring
				maxValue = checkMax(maxValue, i-stack[len(stack)-1])
			} else {
				// if the stack is empty, push the current index as base
				stack = append(stack, i)
			}
		}
	}

	return maxValue
}

// checkMax is an helper function to return the maximum of two integers
func checkMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
