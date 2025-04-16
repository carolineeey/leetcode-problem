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
	//solution.NextPermutation([]int{1, 2, 3})
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
	//fmt.Println(solution.IsValidSudoku(board))
	// 37. Sudoku Solver
	solution.SolveSudoku(board)
	// 39. Combination Sum
	//fmt.Println(solution.CombinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// 40. Combination Sum II
	//fmt.Println(solution.CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// 43. Multiply Strings
	//fmt.Println(solution.Multiply("4", "3"))
	// 44. Wildcard Match
	//fmt.Println(solution.IsMatch("aa", "*"))
	// 45. Jump Game 2
	//fmt.Println(solution.Jump([]int{2, 4, 1, 2, 3, 1, 1, 2}))
	// 46. Permutations
	//fmt.Println(solution.Permute([]int{1, 1, 2}))
	// 47. Permutations
	//fmt.Println(solution.PermuteUnique([]int{1, 1, 2}))
	// 48. Rotate Image
	//solution.Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	// 49. Group Anagrams
	//fmt.Println(solution.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// 50. Pow(x,n)
	//fmt.Println(solution.MyPow(2.0000, 6))
	// 53. Max Subarray
	//fmt.Println(solution.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// 54. Spiral Matrix
	//fmt.Println(solution.SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// 55. Jump Game
	//fmt.Println(solution.CanJump([]int{3, 2, 1, 3, 4}))
	// 56. Merge Intervals
	//fmt.Println(solution.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// 57. Insert Interval
	//fmt.Println(solution.Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	// 58. Length of Last Word
	//fmt.Println(solution.LengthOfLastWord("   fly me   to   the moon  "))
	// 59. Spiral Matrix 2
	//fmt.Println(solution.GenerateMatrix(1))
	// 60. Permutation Sequence
	//fmt.Println(solution.GetPermutation(3, 5))
	// 66. Plus One
	//fmt.Println(solution.PlusOne([]int{9, 9}))
	// 67. Add Binary
	//fmt.Println(solution.AddBinary("101", "101"))
	// 69. Sqrt(x)
	//fmt.Println(solution.MySqrt(1))
	// 70. Climbing Stairs
	//fmt.Println(solution.ClimbStairs(2))
	// 121. Best Time to Buy and Sell Stock
	//fmt.Println("max profit", solution.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	// 122. Best Time to Buy and Sell Stock 2
	//fmt.Println("max profit 2", solution.MaxProfit2([]int{7, 1, 5, 3, 6, 4}))
	// 123. Best Time to Buy and Sell Stock 3
	//fmt.Println("max profit 3", solution.MaxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	// 188. Best Time to Buy and Sell Stock 4
	//fmt.Println("max profit 4", solution.MaxProfit4(2, []int{3, 2, 6, 5, 0, 3}))
	// 242. Valid Anagram
	//fmt.Println(solution.IsAnagramV2("anagram", "nagaram"))
	// 367. Is Perfect Square
	//fmt.Println(solution.IsPerfectSquare(12))
	// 438. Find All Anagrams
	//fmt.Println(solution.FindAnagrams("cbaebabacd", "abc"))
	// 567. Permutation in String
	//fmt.Println(solution.CheckInclusion("ad", "eidbaooo"))
	// 633. Sum of Square Numbers
	//fmt.Println(solution.JudgeSquareSum(26))
	// 989. Add to Array Form
	fmt.Println(solution.AddToArrayForm([]int{1, 2, 0, 0}, 34))
	// 1588. Sum of All Odd Lengths
	//fmt.Println(solution.SumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	// 1886. Determine Whether Matrix Can Be Obtained By Rotation
	//fmt.Println(solution.FindRotation([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}))
	// 2778. Sum of Squares of Special Elements
	//fmt.Println(solution.SumOfSquares([]int{2, 7, 1, 19, 18, 3}))
	// 3264. Final Array State After K Multiplication Operations I
	//fmt.Println(solution.GetFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
}
