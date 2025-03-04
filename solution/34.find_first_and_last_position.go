package solution

func SearchRange(nums []int, target int) []int {
	first := findBound(nums, target, true)
	if first == -1 {
		return []int{-1, -1} //target not found
	}

	last := findBound(nums, target, false)

	return []int{first, last}
}

func findBound(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if isFirst {
				// check if it's first occurrence
				if mid == left || nums[mid-1] != target {
					return mid
				}

				right = mid - 1
			} else {
				// check if it's the last occurrence
				if mid == right || nums[mid+1] != target {
					return mid
				}

				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
