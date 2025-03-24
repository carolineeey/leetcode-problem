package solution

func SumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	result := 0

	for i := 0; i < n; i++ {
		count := ((i+1)*(n-i) + 1) / 2
		result += arr[i] * count
	}

	return result
}
