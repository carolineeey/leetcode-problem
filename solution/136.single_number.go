package solution

func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR (exclusive or) cancels out duplicates
	}

	return result
}
