package solution

func MostFrequentEven(nums []int) int {
	freq := make(map[int]int)
	res := -1
	maxFreq := 0

	for _, num := range nums {
		if num%2 == 0 {
			freq[num]++
			if freq[num] > maxFreq || (freq[num] == maxFreq && num < res) {
				maxFreq = freq[num]
				res = num
			}
		}
	}

	return res
}
