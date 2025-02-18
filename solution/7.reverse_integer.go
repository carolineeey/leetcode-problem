package solution

import "math"

func Reverse(x int) int {
	if x == 0 {
		return 0
	}

	isNegative := x < 0
	if isNegative {
		x = -x
	}

	reversedInt := 0
	for x != 0 {
		// extract the last digit
		digit := x % 10
		// shift left (multiply by 10) and add the digit
		reversedInt = reversedInt*10 + digit
		// remove the last digit from x
		x = x / 10
	}

	// handle overflow (if the reversed integer is out of bounds)
	if reversedInt > math.MaxInt32 {
		return 0
	}

	// if negative, return the negative reversed number
	if isNegative {
		return -reversedInt
	}

	return reversedInt
}
