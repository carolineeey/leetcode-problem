package solution

import (
	"sort"
)

func FindKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
