package solution

func RomanToInt(s string) int {
	// mapping of roman numerals to their corresponding integer values
	valueMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	value := 0

	for i := 0; i < len(s); i++ {
		// get the value of the current roman character
		current := valueMap[s[i]]
		// if current character is smaller than the next character, we subtract it
		if i+1 < len(s) && current < valueMap[s[i+1]] {
			value -= current
		} else {
			value += current
		}
	}

	return value
}
