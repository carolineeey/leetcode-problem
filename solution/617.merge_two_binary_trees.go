package solution

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// If either node is nil, return the non-nil one (or nil if both are nil)
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	merged := &TreeNode{Val: root1.Val + root2.Val}
	merged.Left = MergeTrees(root1.Left, root2.Left)
	merged.Right = MergeTrees(root1.Right, root2.Right)

	return merged
}
