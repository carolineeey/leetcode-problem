package solution

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// treat the digit as a byte (ASCII value) rather than a character.
	letterMap := map[rune][]string{
		50: {"a", "b", "c"}, //The ASCII value of '2' is 50
		51: {"d", "e", "f"}, //The ASCII value of '3' is 51
		52: {"g", "h", "i"},
		53: {"j", "k", "l"},
		54: {"m", "n", "o"},
		55: {"p", "q", "r", "s"},
		56: {"t", "u", "v"},
		57: {"w", "x", "y", "z"},
	}

	result := []string{" "}

	for _, digit := range digits {
		letters := letterMap[digit]
		var temp []string

		// for each existing combination in the result, append each letter from the current digit
		for _, c := range result {
			for _, l := range letters {
				temp = append(temp, c+l)
			}
		}

		// now set result to temp, containing the new combinations.
		result = temp
	}

	return result
}
