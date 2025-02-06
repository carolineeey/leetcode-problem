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
	firstArray := []int{7, 5, 6, 8}
	secondArray := []int{1, 2, 3, 4}
	median := solution.FindMedianSortedArrays(firstArray, secondArray)
	fmt.Println(median)
	//// 9. Palindrome Number
	//isp := solution.IsPalindrome(123211)
	//fmt.Println(isp)
	//

	// 13. Roman to Integer
	//val := solution.RomanToInt("MDCCLXXIV")
	//fmt.Println(val)
	// 17. Letter Combinations of a Phone Number
	//combinations := solution.LetterCombinations("234")
	//fmt.Println(combinations)
	// 20. Valid parentheses
	//isValid := solution.IsValid("()[]{(]}")
	//fmt.Println(isValid)

	//// 26. Remove duplicates from Sorted Array
	//nums = []int{5, 2, 3, 2, 4, 4}
	//length := solution.RemoveDuplicates(nums)
	//fmt.Println(length)

	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//length := solution.RemoveElement(nums, 2)
	//fmt.Println(length)
}
