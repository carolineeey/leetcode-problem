package solution

import (
	"sort"
)

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	} else if len(nums) == 1 && nums[0] == target {
		return []int{target, target}
	}

	sort.Ints(nums)

	result := []int{-1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}

	if len(result) == 1 {
		return []int{-1, -1}
	}

	return result[1:len(result)]
}
