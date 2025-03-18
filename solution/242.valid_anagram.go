package solution

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sortedString := sortString(t)
	sortedAnagram := sortString(s)

	if sortedAnagram != sortedString {
		return false
	}

	return true
}
