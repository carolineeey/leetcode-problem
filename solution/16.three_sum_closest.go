package solution

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	minDifference := math.MaxInt32

	// iterate over each element as the first element of the triplet
	for i := 0; i < len(nums)-2; i++ {
		// skip duplicate element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// initiate the left and right pointer
		l, r := i+1, len(nums)-1
		for l < r {
			// calculate current sum of the triplet
			sum := nums[i] + nums[l] + nums[r]
			// if the sum match to the target, return it immediately
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}

			// calculate the difference to the target
			diffToTarget := target - sum
			// calculate absolute value
			if target-sum < 0 {
				diffToTarget = -(target - sum)
			}

			// if the current sum is closest to the target, update the result
			if diffToTarget < minDifference {
				result = sum
				minDifference = diffToTarget
			}
		}
	}

	return result
}
