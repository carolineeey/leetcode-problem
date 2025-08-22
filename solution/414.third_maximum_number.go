package solution

import "sort"

func ThirdMax(nums []int) int {
	// use map to remove duplicate
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}

	// append unique value to arr slice
	arr := make([]int, 0, len(nums))
	for num := range unique {
		arr = append(arr, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if len(arr) >= 3 {
		return arr[2] // third distinct max
	}

	return arr[0] // max if less than 3
}
