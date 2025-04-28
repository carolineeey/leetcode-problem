package solution

func CanBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	openMin := 0 // minimum number of open parentheses assuming worst case
	openMax := 0 // maximum number of open parentheses assuming best case

	for i := 0; i < len(s); i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				openMin++
				openMax++
			} else {
				openMin--
				openMax--
			}
		} else { // can be either ( or )
			openMin--
			openMax++
		}

		if openMax < 0 { // means it's too many unmatched ) — invalid.
			return false
		}

		if openMin < 0 {
			openMin = 0
		}
	}

	return openMin == 0 // openMin == 0 means it's possible to fully match
}
