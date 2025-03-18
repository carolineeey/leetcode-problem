package solution

import "sort"

func GroupAnagrams(strs []string) [][]string {
	// a map to store the sorted string as key and list of anagrams as value
	matchAnagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		// append the string to the anagram group in the map
		matchAnagrams[sortedStr] = append(matchAnagrams[sortedStr], str)
	}

	// extracting the groups of anagram
	result := make([][]string, 0)
	for _, str := range matchAnagrams {
		result = append(result, str)
	}

	return result
}

// sortString is an helper function to sort the string
// using rune to avoid issues with multibyte character
func sortString(word string) string {
	s := []rune(word)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return string(s)
}
