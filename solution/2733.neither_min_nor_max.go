package solution

import "sort"

func FindNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)

	if nums[0] == nums[len(nums)-1] {
		return -1
	}

	return nums[1]
}
