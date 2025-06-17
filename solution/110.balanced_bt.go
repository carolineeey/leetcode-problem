package solution

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, balanced := checkBalanced(root)

	return balanced
}

func checkBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalanced(root.Left)
	rightHeight, rightBalanced := checkBalanced(root.Right)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return 1 + max(leftHeight, rightHeight), true
}
