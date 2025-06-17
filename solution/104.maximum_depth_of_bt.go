package solution

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightDepth := MaxDepth(root.Right)
	leftDepth := MaxDepth(root.Left)

	return 1 + max(leftDepth, rightDepth)
}
