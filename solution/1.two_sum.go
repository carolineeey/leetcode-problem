package solution

func TwoSumV1(nums []int, target int) (result []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func TwoSumV2(nums []int, target int) (result []int) {
	numIndices := make(map[int]int)
	for i, num := range nums {
		complement := target - num

		if index, found := numIndices[complement]; found {
			result = []int{index, i}
			return
		}

		numIndices[num] = i
	}

	return result
}
