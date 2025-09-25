package solution

import "sort"

func SortedSquares(nums []int) []int {
	var squares []int
	for i := 0; i < len(nums); i++ {
		squares = append(squares, nums[i]*nums[i])
	}
	sort.Ints(squares)
	return squares
}
