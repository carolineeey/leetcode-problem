package solution

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	currentCombination := make([]int, 0)

	sort.Ints(candidates)

	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target == 0 {
			// make a copy of currentCombination to be returned as result
			combination := make([]int, len(currentCombination))
			copy(combination, currentCombination)
			result = append(result, combination)
			return
		}

		for i := start; i < len(candidates); i++ {
			// skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > target {
				// if candidate is greater than target then no need to continue
				break
			}

			// include the candidate in the current combination
			currentCombination = append(currentCombination, candidates[i])
			// backtrack since the same number can be used
			backtrack(i+1, target-candidates[i])
			// remove the last added candidate
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, target)
	return result
}
