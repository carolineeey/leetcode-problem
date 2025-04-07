package solution

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	// sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int

	for _, interval := range intervals {
		// append if result is empty or not overlap
		if len(result) == 0 || result[len(result)-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			// merge intervals
			result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
		}
	}

	return result
}
