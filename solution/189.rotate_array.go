package solution

func RotateArray(nums []int, k int) {
	n := len(nums)
	// if k is greater than the length of the array, rotate only the necessary number of steps.
	k = k % n

	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	// reverse the whole array
	reverse(0, n-1)
	// reverse the first k element
	reverse(0, k-1)
	// reverse the rest (from k to end)
	reverse(k, n-1)
}
