package solution

func IsValid(s string) bool {
	parenthesesMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		// check if the current character is a closing parenthesis
		if opening, ok := parenthesesMap[s[i]]; ok {
			// check if stack is empty or the top of the stack doesn't match the corresponding opening parenthesis
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// pop the top of the stack (match found)
			stack = stack[:len(stack)-1]
		} else {
			// if it's an opening parenthesis, push it onto the stack
			stack = append(stack, s[i])
		}
	}

	// if the stack is empty, all parentheses are matched
	return len(stack) == 0
}
