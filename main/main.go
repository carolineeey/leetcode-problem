package main

import (
	"fmt"
	"github.com/carolineeey/leetcode-problem/solution"
)

func main() {
	////1. Two Sum
	//nums := []int{3, 2, 4}
	//target := 6
	//result := solution.TwoSumV2(nums, target)
	//fmt.Println(result)
	//
	//// 2. Add two numbers
	//l1 := &solution.ListNode{
	//	Val: 4,
	//	Next: &solution.ListNode{
	//		Val: 2,
	//		Next: &solution.ListNode{
	//			Val: 3,
	//		},
	//	},
	//}
	//
	//l2 := &solution.ListNode{
	//	Val: 2,
	//	Next: &solution.ListNode{
	//		Val: 1,
	//		Next: &solution.ListNode{
	//			Val: 5,
	//		},
	//	},
	//}
	//
	//list := solution.AddTwoNumbers(l1, l2)
	//fmt.Println(list)
	// 3. Longest Substring Without Repeating Characters
	//length := solution.LengthOfLongestSubstring("pwwkew")
	//fmt.Println(length)
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
	// 20. Valid parentheses
	//fmt.Println(solution.IsValid("()[]{()}"))
	// 26. Remove duplicates from Sorted Array
	nums := []int{5, 2, 3, 2, 4, 4}
	fmt.Println(solution.RemoveDuplicates(nums))
	// 27. Remove Element
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//length := solution.RemoveElement(nums, 2)
	//fmt.Println(length)
	// 28. Find the Index of the First Occurrence in a String
	//res := solution.StrStr("sadbutsad", "sad")
	//fmt.Println(res)
	// 35. Search Insert Position
	//fmt.Println(solution.SearchInsert([]int{1}, 1))
	// 58. Length of Last Word
	//fmt.Println(solution.LengthOfLastWord("   fly me   to   the moon  "))
	// 70. Climbing Stairs
	//fmt.Println(solution.ClimbStairs(2))
}
