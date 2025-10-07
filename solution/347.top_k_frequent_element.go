package solution

import "sort"

func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Convert map to slice of pairs
	pairs := make([][2]int, 0, len(freq))
	for num, count := range freq {
		pairs = append(pairs, [2]int{num, count})
	}

	// Sort by frequency (descending)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// Pick top k elements
	result := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i][0])
	}

	return result
}
