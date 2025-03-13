package solution

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {
	var res [][]int
	n := len(nums)

	// Sort nums to handle duplicates
	sort.Ints(nums)

	var backtrack func(first int)
	backtrack = func(first int) {
		if first == n {
			tmp := make([]int, n)
			copy(tmp, nums)
			res = append(res, tmp)
		}

		used := make(map[int]bool)
		for i := first; i < n; i++ {
			if used[nums[i]] {
				continue // Skip duplicates
			}
			used[nums[i]] = true

			// Swap the current element with the 'first' element
			nums[first], nums[i] = nums[i], nums[first]

			// Recursively call backtrack with the next index
			backtrack(first + 1)

			// Restore the array to its original state
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)
	return res
}
