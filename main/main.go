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
	//				Next: &solution.ListNode{
	//					Val: 5,
	//				},
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
	//board := [][]byte{
	//	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	//fmt.Println(solution.IsValidSudoku(board))
	// 37. Sudoku Solver
	//solution.SolveSudoku(board)
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
	// 61. Rotate List
	//fmt.Println(solution.RotateRight(l1, 2))
	// 62. Unique Paths
	//fmt.Println(solution.UniquePaths(3, 7))
	// 63. Unique Paths With Obstacle
	//fmt.Println(solution.UniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	// 64. Minimum Path Sum
	//fmt.Println(solution.MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	// 66. Plus One
	//fmt.Println(solution.PlusOne([]int{9, 9}))
	// 67. Add Binary
	//fmt.Println(solution.AddBinary("101", "101"))
	// 68. Text Justification
	//fmt.Println(solution.FullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	// 69. Sqrt(x)
	//fmt.Println(solution.MySqrt(1))
	// 70. Climbing Stairs
	//fmt.Println(solution.ClimbStairs(2))
	// 71. Simplify Path
	//fmt.Println(solution.SimplifyPath("/home//foo/"))
	// 73. Set Matrix Zeroes
	//solution.SetZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	// 74. Search in 2D Matrix
	//fmt.Println(solution.SearchMatrixWithBinary([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	// 75. Sort Colors
	//solution.SortColors([]int{2, 0, 2, 1, 1, 0})
	// 77. Combine
	//fmt.Println(solution.Combine(4, 2))
	// 78. Subsets
	//fmt.Println(solution.Subsets([]int{1, 2, 3}))
	// 79. Word Search
	//fmt.Println(solution.Exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	// 80. Remove Duplicates from Sorted Array
	//fmt.Println(solution.RemoveDuplicates2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	// 81. Search in Rotated Array 2
	//fmt.Println(solution.Search2([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	// 82. Remove Duplicates
	//fmt.Println(solution.DeleteDuplicates2(l1))
	// 83. Remove Duplicate
	//fmt.Println(solution.DeleteDuplicates(l1))
	// 88. Merge Sorted Array
	//solution.MergeSorted([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	// 84. Largest Rectangle in Histogram
	//fmt.Println(solution.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	// 85. Maximal Rectangle Area
	//fmt.Println(solution.MaximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	p := &solution.TreeNode{
		Val: 1,
		Left: &solution.TreeNode{
			Val: 2,
			Left: &solution.TreeNode{
				Val: 3,
			},
			Right: &solution.TreeNode{
				Val: 4,
			},
		},
		Right: &solution.TreeNode{
			Val: 5,
		},
	}
	q := &solution.TreeNode{
		Val: 1,
		Left: &solution.TreeNode{
			Val: 2,
			Left: &solution.TreeNode{
				Val: 3,
			},
			Right: &solution.TreeNode{
				Val: 2,
			},
		},
		Right: &solution.TreeNode{
			Val: 5,
		},
	}
	r := &solution.TreeNode{
		Val: 10,
		Left: &solution.TreeNode{
			Val: 2,
		},
		Right: &solution.TreeNode{
			Val: 7,
		},
	}
	s := &solution.TreeNode{
		Val: 10,
		Left: &solution.TreeNode{
			Val: 5,
			Left: &solution.TreeNode{
				Val: 3,
			},
			Right: &solution.TreeNode{
				Val: 7,
			},
		},
		Right: &solution.TreeNode{
			Val: 15,
			Right: &solution.TreeNode{
				Val: 18,
			},
		},
	}

	t := &solution.Node{
		Val: 10,
		Left: &solution.Node{
			Val: 5,
			Left: &solution.Node{
				Val: 3,
			},
			Right: &solution.Node{
				Val: 7,
			},
		},
		Right: &solution.Node{
			Val: 15,
			Right: &solution.Node{
				Val: 18,
			},
		},
		Next: &solution.Node{
			Val: 9,
			Right: &solution.Node{
				Val: 5,
			},
		},
	}

	// 90. Subsets 2
	//fmt.Println(solution.SubsetsWithDup([]int{1, 2, 2}))
	// 91. Decode Ways
	//fmt.Println(solution.NumDecoding("226"))
	// 92. Reverse Linked Link 2
	head := &solution.ListNode{
		Val: 1,
		Next: &solution.ListNode{
			Val: 2,
			Next: &solution.ListNode{
				Val: 3,
				Next: &solution.ListNode{
					Val: 4,
					Next: &solution.ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	//fmt.Println(solution.ReverseBetween(head, 2, 4))
	// 93. Restore IP Address
	//fmt.Println(solution.RestoreIpAddresses("25525511135"))
	// 94. Binary Tree Inorder Traversal
	fmt.Println(solution.InorderTraversal(p))
	// 95. Update Binary Search Tree (BST)
	solution.GenerateTrees(3)
	// 96. Unique Binary Search Tree (BST)
	solution.NumTrees(3)
	// 98. Validate Binary Search Tree
	//fmt.Println(solution.IsValidBST(q))
	// 99. Recover Binary Search Tree
	solution.RecoverTree(p)
	// 100. Tree Node
	solution.IsSameTree(p, q)
	// 101. Symmetric Tree
	fmt.Println("101. Symmetric Tree", solution.IsSymmetric(p))
	// 102. Binary Tree Level Order Traversal
	fmt.Println("102. Binary Tree Level Order Traversal", solution.LevelOrder(p))
	// 104. Maximum Depth of Binary Tree
	fmt.Println("104. Maximum Depth of Binary Tree", solution.MaxDepth(p))
	// 105. Construct Binary Tree from Preorder and Inorder Traversal
	fmt.Println("105. Construct Binary Tree from Preorder and Inorder Traversal", solution.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	// 106. Construct Binary Tree from Inorder and Postorder Traversal
	fmt.Println("106. Construct Binary Tree from Inorder and Postorder Traversal", solution.BuildTreeTraversal([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	// 107. Binary Tree Level Order Traversal II
	fmt.Println("107. Binary Tree Level Order Traversal II", solution.LevelOrderBottom(p))
	// 108. Convert Sorted Array to Binary Search Tree
	fmt.Println("108. Convert Sorted Array to Binary Search Tree", solution.SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	// 109. Convert Sorted List to Binary Search Tree
	fmt.Println("109. Convert Sorted List to Binary Search Tree", solution.SortedListToBST(head))
	// 110. Balanced Binary Tree
	fmt.Println("110. Balanced Binary Tree", solution.IsBalanced(q))
	// 111. Minimum Byte of The Binary Tree
	fmt.Println("111. Minimum Byte of The Binary Tree", solution.MinDepth(p))
	// 112. Path Sum
	fmt.Println("112. Path Sum", solution.HasPathSum(r, 9))
	// 113. Path Sum 2
	fmt.Println("113. Path Sum", solution.PathSum(r, 1))
	// 114. Flatten Binary Tree to Linked List
	solution.Flatten(r)
	// 116. Populating Next Right Pointers in Each Node
	fmt.Println("116. Populating Next Right Pointers in Each Node", solution.Connect(t))
	// 117. Populating Next Right Pointers in Each Node
	fmt.Println("117. Populating Next Right Pointers in Each Node 2", solution.Connect2(t))
	// 118. Pascal Triangle
	fmt.Println("118. Pascal Triangle", solution.Generate(5))
	// 119. Pascal Triangle 2
	fmt.Println("119. Pascal Triangle 2", solution.GeneratePascalTriangle(4))
	// 120. Triangle
	fmt.Println("120. Triangle", solution.MinimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	// 121. Best Time to Buy and Sell Stock
	//fmt.Println("max profit", solution.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	// 122. Best Time to Buy and Sell Stock 2
	//fmt.Println("max profit 2", solution.MaxProfit2([]int{7, 1, 5, 3, 6, 4}))
	// 123. Best Time to Buy and Sell Stock 3
	//fmt.Println("max profit 3", solution.MaxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	// 124. Binary Tree Max Path Sum
	fmt.Println("124. Binary Tree Max Path Sum", solution.MaxPathSum(s))
	// 125. Valid Palindrome
	//fmt.Println("125. Valid Palindrome 1", solution.IsCharacterPalindrome("A man, a plan, a canal: Panama"))
	// 128. Longest Consecutive Sequence
	fmt.Println("128. Longest Consecutive Sequence", solution.LongestConsecutive2([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// 129. Sum Root to Leaf Numbers
	fmt.Println("129. Sum Root to Leaf Numbers", solution.SumNumbers(r))
	// 136. Single Number
	fmt.Println("136. Single Number", solution.SingleNumber([]int{4, 1, 2, 1, 2}))
	// 137. Single Number 2
	fmt.Println("137. Single Number 2", solution.SingleNumber2([]int{4, 1, 1, 2, 2, 1, 2}))
	// 139. Word Break
	fmt.Println("139. Word Break", solution.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// 141. Linked List Cycle
	fmt.Println("141. Linked List Cycle", solution.HasCycle(head))
	// 142. Linked List Cycle 2
	fmt.Println("142. Linked List Cycle", solution.DetectCycle(head))
	// 168. Excel Sheet Column Title
	fmt.Println("168. Excel Sheet Column Title", solution.ConvertToTitle(701))
	// 171. Excel Sheet Column Number
	fmt.Println("171. Excel Sheet Column Number", solution.TitleToNumber("AB"))
	// 188. Best Time to Buy and Sell Stock 4
	//fmt.Println("max profit 4", solution.MaxProfit4(2, []int{3, 2, 6, 5, 0, 3}))
	// 189. Rotate Array
	//solution.RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	// 202. Happy Number
	fmt.Println("202. Happy Number", solution.IsHappy(19))
	// 204. Count Primes
	fmt.Println("204. Count Primes", solution.CountPrimes(10))
	// 206. Reverse Linked List
	//fmt.Println(solution.ReverseList(head))
	// 215. Kth Largest Element in an Array
	//fmt.Println("215. Kth Largest Element in an Array", solution.FindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	// 221. Maximal Square
	//fmt.Println(solution.MaximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	// 234. Palindrome Linked List
	//solution.IsPalindromeLinkedList(head)
	// 240. Search 2D Matrix 2
	//fmt.Println(solution.SearchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	// 242. Valid Anagram
	//fmt.Println(solution.IsAnagramV2("anagram", "nagaram"))
	// 257. Binary Tree Paths
	//fmt.Println("257. Binary Tree Paths", solution.BinaryTreePaths(p))
	// 258. Add Digits
	//fmt.Println("258. Add Digits", solution.AddDigits(38))
	// 260. Single Number 3
	//fmt.Println("260. Single Number 3", solution.SingleNumber3([]int{1, 2, 1, 3, 2, 5}))
	// 263. Ugly Number
	//fmt.Println("263. Ugly Number", solution.IsUgly(14))
	// 264. Ugly Number 2
	//fmt.Println("264. Ugly Number 2", solution.NthUglyNumber(10))
	// 268. Missing Number
	//fmt.Println("268. Missing Number", solution.MissingNumber([]int{0, 1}))
	// 279. Perfect Squares
	fmt.Println(solution.NumSquares(13))
	// 287. Find the Duplicate Number
	//fmt.Println("287. Find the Duplicate Number", solution.FindDuplicate([]int{1, 3, 4, 2, 2}))
	// 367. Is Perfect Square
	//fmt.Println(solution.IsPerfectSquare(12))
	// 389. Find The Difference
	//fmt.Println("389. Find The Difference", solution.FindTheDifference("abcd", "abcde"))
	// 414. Third Maximum Number
	//fmt.Println("414. Third Maximum Number", solution.ThirdMax([]int{3, 2, 2, 1}))
	// 438. Find All Anagrams
	//fmt.Println(solution.FindAnagrams("cbaebabacd", "abc"))
	// 496. Next Greater Element
	//fmt.Println("496. Next Greater Element", solution.NextGreaterElement([]int{4, 1, 2}, []int{4, 3, 1, 2}))
	// 501. Find Mode in Binary Search Tree
	//fmt.Println(solution.FindMode(p))
	// 515. Find Largest Value in Each Tree Row
	//fmt.Println("515. Find Largest Value in Each Tree Row", solution.LargestValues(p))
	// 567. Permutation in String
	//fmt.Println(solution.CheckInclusion("ad", "eidbaooo"))
	// 627. Merge Two Binary Trees
	//fmt.Println("627. Merge Two Binary Trees", solution.MergeTrees(p, q))
	// 633. Sum of Square Numbers
	//fmt.Println(solution.JudgeSquareSum(26))
	// 637. Average of Levels in Binary Tree
	//fmt.Println("637. Average of Levels in Binary Tree", solution.AverageOfLevels(p))
	// 645. Set Mismatch
	//fmt.Println("645. Set Mismatch", solution.FindErrorNums([]int{1, 2, 2, 4}))
	// 680. Valid Palindrome 2
	//fmt.Println("680. Valid Palindrome 2", solution.ValidPalindrome("abdb_a"))
	// 929. Unique Email Addresses
	//fmt.Println("929. Unique Email Addresses", solution.NomUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	// 938. Range Sum of BST
	//fmt.Println("938. Range Sum of BST", solution.RangeSumBST(s, 7, 15))
	// 965. Univalued Binary Tree
	//fmt.Println("965. Univalued Binary Tree", solution.IsUnivalTree(p))
	// 989. Add to Array Form
	//fmt.Println(solution.AddToArrayForm([]int{1, 2, 0, 0}, 34))
	// 1005. Maximize Sum Of Array After K Negations
	//fmt.Println("1005. Maximize Sum Of Array After K Negations", solution.LargestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	// 1030. Matrix Cells in Distance Order
	//fmt.Println("1030. Matrix Cells in Distance Order", solution.AllCellsDistOrder(1, 2, 0, 0))
	// 1037. Valid Boomerang
	fmt.Println("1037. Valid Boomerang", solution.IsBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// 1232. Check If It Is a Straight Line
	//fmt.Println("1232. Check If It Is a Straight Line", solution.CheckStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	// 1356. Sort Integers by The Number of 1 Bits
	// 1351. Count Negative Numbers in a Sorted Matrix
	fmt.Println("1351. Count Negative Numbers in a Sorted Matrix", solution.CountNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	//fmt.Println("1356. Sort Integers by The Number of 1 Bits", solution.SortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	// 1588. Sum of All Odd Lengths
	//fmt.Println(solution.SumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	// 1886. Determine Whether Matrix Can Be Obtained By Rotation
	//fmt.Println(solution.FindRotation([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}))
	// 2099.Find Subsequence of Length K With the Largest Sum
	//fmt.Println("2099.Find Subsequence of Length K With the Largest Sum", solution.MaxSubsequence([]int{2, 1, 3, 3}, 2))
	// 2116. Check If a Parentheses String Can Be Valid
	//fmt.Println(solution.CanBeValid("))()))", "010100"))
	// 2194. Cells in a Range on an Excel Sheet
	//fmt.Println("2194. Cells in a Range on an Excel Sheet", solution.CellsInRange("K1:L2"))
	// 2236. Root Equals Sum of Children
	//fmt.Println("2236. Root Equals Sum of Children", solution.CheckTree(r))
	// 2427. Number of Common Factors
	//fmt.Println("2427. Number of Common Factors", solution.CommonFactors(12, 6))
	// 2529. Maximum Count of Positive Integer and Negative Integer
	fmt.Println("2529. Maximum Count of Positive Integer and Negative Integer", solution.MaximumCount([]int{-3, -2, -1, 0, 0, 1, 2}))
	// 2708. Maximum Strength Of a Group
	//fmt.Println(solution.MaxStrength([]int{3, -1, -5, 2, 5, -9}))
	// 2733. Find Neither Minimum nor Maximum
	//fmt.Println("2733. Find Neither Minimum nor Maximum", solution.FindNonMinOrMax([]int{3, 2, 1, 4}))
	// 2778. Sum of Squares of Special Elements
	//fmt.Println(solution.SumOfSquares([]int{2, 7, 1, 19, 18, 3}))
	// 3158. Find the XOR of Numbers Which Appear Twice
	//fmt.Println("3158. Find the XOR of Numbers Which Appear Twice", solution.DuplicateNumbersXOR([]int{1, 2, 1, 2}))
	// 3264. Final Array State After K Multiplication Operations I
	//fmt.Println(solution.GetFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
}
