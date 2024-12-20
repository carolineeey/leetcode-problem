package main

import (
	"fmt"
	"github.com/carolineeey/leetcode-problem/solution"
)

func main() {
	// Two Sum
	nums := []int{3, 2, 4}
	target := 6
	result := solution.TwoSumV2(nums, target)
	fmt.Println(result)

	// Palindrome Number
	isp := solution.IsPalindrome(123211)
	fmt.Println(isp)
}
