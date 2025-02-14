package solution

func SearchInsert(nums []int, target int) int {
	// use binary search to find the correct position for the target.
	left, right := 0, len(nums)-1

	// binary search loop
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid // target found, return its index
		} else if nums[mid] < target {
			left = mid + 1 // target is greater, search right half
		} else {
			right = mid - 1 // target is smaller, search left half
		}
	}

	// if target is not found, 'left' will be the insert position
	return left
}
