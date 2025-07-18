package solution

func IsUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	return IsUnivalTree(root.Left) && IsUnivalTree(root.Right)
}
