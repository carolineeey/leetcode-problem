package solution

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar, minSoFar := nums[0], nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]

		// swap if negative
		if n < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		maxSoFar = max(n, maxSoFar*n)
		minSoFar = min(n, minSoFar*n)

		result = max(result, maxSoFar)
	}

	return result
}
