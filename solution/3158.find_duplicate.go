package solution

func DuplicateNumbersXOR(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	answer := 0
	for x, i := range freq {
		// If it's the second time seeing it, XOR with ans
		if i == 2 {
			answer ^= x
		}
	}

	return answer
}
