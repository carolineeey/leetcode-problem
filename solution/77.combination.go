package solution

func Combine(n int, k int) [][]int {
	var res [][]int

	//start: the current number to start from (to avoid duplicates)
	//comb: the current combination being built
	var backtrack func(start int, comb []int)

	backtrack = func(start int, comb []int) {
		if len(comb) == k {
			// make a copy of comb before appending
			tmp := make([]int, k)
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}

		// Prune the loop to avoid unnecessary calls
		for i := start; i <= n-(k-len(comb))+1; i++ {
			backtrack(i+1, append(comb, i))
		}
	}

	backtrack(1, []int{})

	return res
}
