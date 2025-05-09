package solution

import "strings"

func LengthOfLastWord(s string) int {
	ptr := len(s) - 1
	length := 0

	// skip spaces
	for ptr >= 0 && s[ptr] == ' ' {
		ptr--
	}

	// count the length of the last word
	for ptr >= 0 && s[ptr] != ' ' {
		length++
		ptr--
	}

	return length
}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)

	return len(words[len(words)-1])
}
