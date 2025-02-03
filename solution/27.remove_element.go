package solution

func RemoveElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		// If the element is not equal to val, move it to the 'k' position
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
