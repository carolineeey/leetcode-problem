package solution

// Floyd's algorithm to find the entry point of the cycle — which is the duplicate
func FindDuplicate(nums []int) int {
	// find intersection point
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// find the duplicate
	slow = nums[0]
	if slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
