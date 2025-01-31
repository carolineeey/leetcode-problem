package solution

import (
	"sort"
)

func RemoveDuplicates(nums []int) int {
	sort.Ints(nums)
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l += 1
		}
	}
	return l
}
