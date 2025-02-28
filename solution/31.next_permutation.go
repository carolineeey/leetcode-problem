package solution

func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// find the first decreasing element from the end
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// find the element that larger than nums[i] to swap with
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}

		// swap elements at i and j
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
