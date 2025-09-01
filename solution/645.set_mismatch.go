package solution

func FindErrorNums(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)
	duplicate, missing := -1, -1

	for _, num := range nums {
		count[num]++
	}

	for i := 1; i <= n; i++ {
		if count[i] == 2 {
			duplicate = i
		} else if count[i] == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}
