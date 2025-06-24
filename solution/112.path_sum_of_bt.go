package solution

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// check if it is a leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// recur left and right with updated target
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}
