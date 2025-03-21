package solution

import (
	"math"
)

func JudgeSquareSum(c int) bool {
	// use two pointers approach
	// get the square of c first then check if there exist two integers
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}

	return false
}
