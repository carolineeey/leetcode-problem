package main

import (
	"fmt"
	"github.com/carolineeey/leetcode-problem/solution"
)

func main() {
	//l1 := &solution.ListNode{
	//	Val: 1,
	//	Next: &solution.ListNode{
	//		Val: 2,
	//		Next: &solution.ListNode{
	//			Val: 3,
	//			Next: &solution.ListNode{
	//				Val: 4,
	//			},
	//		},
	//	},
	//}
	//l2 := &solution.ListNode{
	//	Val: 5,
	//	Next: &solution.ListNode{
	//		Val: 7,
	//		Next: &solution.ListNode{
	//			Val: 2,
	//			Next: &solution.ListNode{
	//				Val: 6,
	//			},
	//		},
	//	},
	//}
	//1. Two Sum
	//fmt.Println(solution.TwoSumV2([]int{3, 2, 4}, 6))
	//
	//// 2. Add two numbers
	//fmt.Println("AddTwoNumbers", solution.AddTwoNumbers(l1, l2))
	// 3. Longest Substring Without Repeating Characters
	//fmt.Println(solution.LengthOfLongestSubstring("pwwkew"))
	// 4. Median of Two Sorted Arrays
	//firstArray := []int{7, 5, 6, 8}
	//secondArray := []int{1, 2, 3, 4}
	//median := solution.FindMedianSortedArrays(firstArray, secondArray)
	//fmt.Println(median)
	// 5. Longest Palindromic Substring
	//fmt.Println(solution.LongestPalindrome("abaabsa"))
	// 6. Zigzag conversion
	//fmt.Println(solution.Convert("PAYPALISHITING", 4))
	// 7. Reverse Integer
	//fmt.Println(solution.Reverse(-321))
	// 8. String to Integer (atoi)
	//fmt.Println(solution.MyAtoi("-010234"))
	// 9. Palindrome Number
	//fmt.Println(solution.IsPalindrome(12321))
	// 11. Container with Most Water
	//fmt.Println(solution.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	// 12. Integer to Roman
	//fmt.Println(solution.IntToRoman(1206))
	// 13. Roman to Integer
	//fmt.Println(solution.RomanToInt("MDCCLXXV"))
	// 14. Longest Common Prefix
	//fmt.Println(solution.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	// 15. 3Sum
	//fmt.Println(solution.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	// 16. 3Sum Closest
	//fmt.Println(solution.ThreeSumClosest([]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, 1))
	// 17. Letter Combinations of a Phone Number
	//fmt.Println(solution.LetterCombinations("234"))
	// 18. 4Sum
	//fmt.Println(solution.FourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	// 19. Remove Nth Node From End of List
	//fmt.Println(solution.RemoveNthFromEnd(l, 2))
	// 20. Valid parentheses
	//fmt.Println(solution.IsValid("()[]{()}"))
	// 21. Merge Two Sorted Lists
	//fmt.Println(solution.MergeTwoLists(l1, l2))
	// 22. Generate Parentheses
	//fmt.Println(solution.GenerateParenthesis(4))
	// 23. Merge k Sorted Lists
	//fmt.Println(solution.MergeKLists([]*solution.ListNode{l1, l2}))
	// 24. Swap Nodes in Pairs
	//fmt.Println(solution.SwapPairs(l1))
	// 26. Remove duplicates from Sorted Array
	//fmt.Println(solution.RemoveDuplicates([]int{5, 2, 3, 2, 1, 4, 4}))
	// 27. Remove Element
	//fmt.Println(solution.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 8, 2, 2}, 2))
	// 28. Find the Index of the First Occurrence in a String
	//fmt.Println(solution.StrStr("sadbutsad", "but"))
	// 29. Divide Two Integers
	//fmt.Println(solution.Divide(10, 3))
	// 31. Next Permutation
	solution.NextPermutation([]int{1, 2, 3})
	// 32. Longest Valid Parentheses
	//fmt.Println(solution.LongestValidParentheses(")()()("))
	// 33. Search in Rotated Sorted Array
	//fmt.Println(solution.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// 34. Find First and Last Position of Element in Sorted Array
	//fmt.Println(solution.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	// 35. Search Insert Position
	//fmt.Println(solution.SearchInsert([]int{1}, 1))
	// 36. Valid Sudoku
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(solution.IsValidSudoku(board))
	// 37. Sudoku Solver
	solution.SolveSudoku(board)
	// 39. Combination Sum
	fmt.Println(solution.CombinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// 40. Combination Sum II
	fmt.Println()
	fmt.Println(solution.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// 43. Multiply Strings
	fmt.Println(solution.Multiply("4", "3"))
	// 44. Wildcard Match
	fmt.Println(solution.IsMatch("aa", "*"))
	// 45. Jump Game 2
	fmt.Println(solution.Jump([]int{2, 4, 1, 2, 3, 1, 1, 2}))
	// 58. Length of Last Word
	//fmt.Println(solution.LengthOfLastWord("   fly me   to   the moon  "))
	// 70. Climbing Stairs
	//fmt.Println(solution.ClimbStairs(2))
}
