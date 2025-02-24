package solution

import (
	"sort"
)

// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	var result [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// skip duplicate value for first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// skip duplicate value for second element
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			l := len(nums) - 1

			// approach the third and fourth element
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					// skip duplicate value for third element
					for k < l && nums[k] == nums[k+1] {
						k++
					}

					// skip duplicate value for fourth element
					for k < l && nums[l] == nums[l-1] {
						l--
					}

					k++
					l--
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}

	return result
}
