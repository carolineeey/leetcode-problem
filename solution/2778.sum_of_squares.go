package solution

func SumOfSquares(nums []int) int {
	n := len(nums)
	result := 0
	// temporary array with 0 as first element to maintain the indexing so it start from 1
	temp := []int{0}
	for i := 0; i < n; i++ {
		temp = append(temp, nums[i])
	}

	for i := 1; i < len(temp); i++ {
		if n%i == 0 {
			result += temp[i] * temp[i]
		}
	}

	return result
}
