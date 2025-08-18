package solution

import "sort"

func MaxSubsequence(nums []int, k int) []int {
	n := len(nums)
	pairs := make([][2]int, n) // value, index
	for i, v := range nums {
		pairs[i] = [2]int{v, i}
	}

	// sort by value descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	// take top k
	topK := pairs[:k]

	// store by index ascending (restore order)
	sort.Slice(topK, func(i, j int) bool {
		return topK[i][1] < topK[j][1]
	})

	// collect value
	res := make([]int, k)
	for i, pair := range topK {
		res[i] = pair[0]
	}

	return res
}
