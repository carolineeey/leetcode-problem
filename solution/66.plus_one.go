package solution

func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	// If a digit becomes 10, set it to 0 and continue carry
	// If all digits become 0 (e.g. [9,9,9]), then prepend a 1 at the front
	return append([]int{1}, digits...)
}
