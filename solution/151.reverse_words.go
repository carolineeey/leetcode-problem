package solution

import (
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)

	// construct reverse words
	//reversed := make([]string, len(words))
	//for i, word := range words {
	//	word = strings.TrimSpace(word)
	//	reversed[len(words)-i-1] = word
	//}

	// reverse words using two-pointer for space efficient
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
