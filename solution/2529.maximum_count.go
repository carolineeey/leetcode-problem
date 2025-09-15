package solution

func MaximumCount(nums []int) int {
	pos, neg := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if nums[i] < 0 {
				neg++
			} else {
				pos++
			}
		}
	}
	return max(pos, neg)
}
