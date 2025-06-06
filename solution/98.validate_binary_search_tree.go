package solution

func IsValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
	// a nil node means you've reached the end of a branch,
	// and an empty tree (or subtree) is always a valid BST
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}

	if max != nil && node.Val >= *max {
		return false
	}

	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}
