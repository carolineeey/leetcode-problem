package solution

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	var res [][]int
	// sort for detecting and skipping duplicates efficiently in the next steps
	sort.Ints(nums)

	// start: index from where to consider elements in nums
	// path: current subset being constructed
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		// make a copy of the current subset and add it to the result
		// to avoid future modifications affecting the stored result
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			// skip duplicates at the same recursive level
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// add nums[i] to the current subset
			path = append(path, nums[i])
			// recurse with the next index (i+1) and the updated subset (path)
			backtrack(i+1, path)
			// remove the last element to explore the next option in the loop
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return res
}
