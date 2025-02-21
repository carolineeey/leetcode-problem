package solution

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var result [][]int

	// sort the number
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// skip if the i is duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// initiate the left and right pointer
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})

				// skip duplicates for left and right
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				// move the pointer if the matching triplet found
				l++
				r--
			} else if sum < 0 {
				// increase if the sum is too small
				l++
			} else {
				// decrease if the sum is too large
				r--
			}
		}
	}
	return result
}
