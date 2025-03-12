package solution

func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	coverage := 0

	for i := 0; i < len(nums); i++ {
		if i > coverage {
			return false
		}
		coverage = max(coverage, i+nums[i])

		if coverage >= len(nums)-1 {
			return true
		}
	}

	return false
}
