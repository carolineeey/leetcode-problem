package main

import (
	"fmt"
	"github.com/carolineeey/leetcode-problem/solution"
)

func main() {
	//1. Two Sum
	nums := []int{3, 2, 4}
	target := 6
	result := solution.TwoSumV2(nums, target)
	fmt.Println(result)

	// 9. Palindrome Number
	isp := solution.IsPalindrome(123211)
	fmt.Println(isp)

	// 26. Remove duplicates from Sorted Array
	nums = []int{5, 2, 3, 2, 4, 4}
	length := solution.RemoveDuplicates(nums)
	fmt.Println(length)

	// 2. Add two numbers
	l1 := &solution.ListNode{
		Val: 4,
		Next: &solution.ListNode{
			Val: 2,
			Next: &solution.ListNode{
				Val: 3,
			},
		},
	}

	l2 := &solution.ListNode{
		Val: 2,
		Next: &solution.ListNode{
			Val: 1,
			Next: &solution.ListNode{
				Val: 5,
			},
		},
	}

	list := solution.AddTwoNumbers(l1, l2)
	fmt.Println(list)
}
