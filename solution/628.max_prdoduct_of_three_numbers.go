package solution

import "sort"

//The maximum product of three numbers can only come from:

//The three largest numbers (e.g. a * b * c where they are the top 3 positives).

// The two smallest numbers (possibly negatives) and the largest number (because (-x) * (-y) * z could be huge).
func MaximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	// case 1 : three largest numbers
	option1 := nums[n-1] * nums[n-2] * nums[n-3]

	// case 2 : two smallest numbers and the largest number
	option2 := nums[0] * nums[1] * nums[n-1]

	if option1 > option2 {
		return option1
	}

	return option2
}
