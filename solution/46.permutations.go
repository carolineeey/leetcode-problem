package solution

func Permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(first int)
	n := len(nums)
	backtrack = func(first int) {
		if first == n {
			temp := make([]int, n)
			copy(temp, nums)
			res = append(res, temp)
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
}
