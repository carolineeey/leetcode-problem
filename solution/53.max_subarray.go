package solution

func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 { // handle edge cases
		return nums[0]
	}

	maxSum := nums[0] // start with first index
	currentSum := nums[0]
	for i := 1; i < n; i++ {
		// choose between extending the current subarray or starting fresh
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
