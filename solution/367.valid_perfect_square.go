package solution

func IsPerfectSquare(num int) bool {
	left, right := 1, num
	mid := 0

	// using binary search
	for left <= right {
		mid = (left + right) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num { // move forward if less than num
			left = mid + 1
		} else { // move backward if greater than num
			right = mid - 1
		}
	}

	return mid*mid == num
}
