package solution

func MySqrt(x int) int {
	left, right := 1, x
	// using binary search
	for left <= right {
		mid := (right + left) / 2
		if mid*mid > x { // move backward
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1 // move forward
		} else {
			return mid
		}
	}

	// return right because we want to round down the square root and return the smaller
	return right
}
