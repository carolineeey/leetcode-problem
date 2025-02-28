package solution

import (
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Divide(dividend int, divisor int) int {
	// handle overflow case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// check the sign for the result
	isNegative := (dividend < 0) != (divisor < 0)

	// work with positive value
	dividend = abs(dividend)
	divisor = abs(divisor)

	result := 0
	for dividend >= divisor {
		tempDivisor, numDivisor := divisor, 1
		for dividend >= tempDivisor {
			dividend -= tempDivisor
			result += numDivisor

			// double the divisor and the number of divisor by using bit shifting
			numDivisor <<= 1
			tempDivisor <<= 1
		}
	}

	if isNegative {
		result = -result
	}

	return result
}
