package solution

func Subsets(nums []int) [][]int {
	var res [][]int
	var curr []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// make a copy of the current subset and add it to the result
		// it is crucial because curr will change during recursion
		subset := make([]int, len(curr))
		copy(subset, curr)
		// appends the copied subset to the final result res (add this version of the current subset to the result)
		res = append(res, subset)

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			// remove the last added element from curr so we can explore other combinations without that element in subsequent iterations
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return res
}
