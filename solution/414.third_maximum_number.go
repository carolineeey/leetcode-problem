package solution

import "sort"

func ThirdMax(nums []int) int {
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}

	arr := make([]int, 0, len(nums))
	for num := range unique {
		arr = append(arr, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if len(arr) >= 3 {
		return arr[2]
	}

	return arr[0]
}
