package solution

func MajorityElement2(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	var result []int
	threshold := len(nums) / 3
	for num, count := range freq {
		if count > threshold {
			result = append(result, num)
		}
	}
	return result
}
