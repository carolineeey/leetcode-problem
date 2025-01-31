package solution

import (
	"sort"
)

func RemoveDuplicates(nums []int) int {
	// sort in non-decreasing order
	sort.Ints(nums)
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[l] = nums[i]
			// only increase l if the nums[i] is different with nums[i-1]
			l += 1
		}
	}

	return l
}
