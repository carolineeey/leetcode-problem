package solution

import "sort"

func LargestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	// flip negative
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	sort.Ints(nums)

	// if k is odd, flip the smallest element
	if k%2 == 1 {
		nums[0] = -nums[0]
	}

	// sum all nums
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
