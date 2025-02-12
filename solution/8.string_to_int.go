package solution

import (
	"math"
	"unicode"
)

func MyAtoi(s string) (result int) {
	sign := 1
	var i int
	n := len(s)

	// skip leading whitespaces
	for i < n && unicode.IsSpace(rune(s[i])) {
		i++
	}

	// check for optional sign
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	// convert digit to integer
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// check the overflow and underflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
