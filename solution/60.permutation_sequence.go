package solution

import (
	"strconv"
)

func GetPermutation(n int, k int) string {
	// precompute factorials
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	// initialize the list of the number from 1 to n
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	k-- // convert to 0 based index
	result := ""
	// determine the k-th permutation using factorial number system
	for i := n; i >= 1; i-- {
		// find the index of the next number to add to the result
		idx := k / factorial[i-1]
		// append the number at the computed index to the result
		result += strconv.Itoa(nums[idx])
		// remove the used number from the list of available numbers
		nums = append(nums[:idx], nums[idx+1:]...)
		// update k to determine the next number
		k %= factorial[i-1]
	}

	return result
}
