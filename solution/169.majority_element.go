package solution

func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] > len(nums)/2 {
			return num
		}
	}

	return -1
}
