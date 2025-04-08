package solution

func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	i := 0
	n := len(intervals)

	// add all intervals that ends before the new interval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// merge overlapping intervals with newInterval
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
