package solution

func RemoveDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	// start from index 2 because the first two elements are always allowed
	length := 2

	for i := 2; i < n; i++ {
		if nums[i] != nums[length-2] {
			// compare current with the element two positions before
			nums[length] = nums[i]
			length++
		}
	}

	return length
}
