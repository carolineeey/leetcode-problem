package solution

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Pick middle element to be used as the root
	// This ensures the tree stays balanced, left part of the array will form the left subtree,
	// right part will form the right subtree
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])

	return root
}
