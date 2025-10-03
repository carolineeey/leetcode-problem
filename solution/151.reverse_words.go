package solution

import (
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)

	// construct reverse words
	reversed := make([]string, len(words))
	for i, word := range words {
		word = strings.TrimSpace(word)
		reversed[len(words)-i-1] = word
	}

	return strings.Join(reversed, " ")
}
