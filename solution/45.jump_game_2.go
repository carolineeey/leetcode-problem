package solution

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	totalJump := 0
	destination := len(nums) - 1
	coverage, lastJumpIdx := 0, 0

	for i := 0; i < len(nums); i++ {
		coverage = max(coverage, i+nums[i])

		if i == lastJumpIdx {
			lastJumpIdx = coverage
			totalJump++

			if coverage >= destination {
				return totalJump
			}
		}
	}

	return totalJump
}
