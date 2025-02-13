package solution

import (
	"sort"
)

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	// compare characters of the first and last strings
	first := strs[0]
	last := strs[length-1]
	i := 0
	for i < len(first) && i < len(last) && first[i] == last[i] {
		i++
	}

	// the common prefix is the substring from the start to the matched length
	return first[:i]
}
