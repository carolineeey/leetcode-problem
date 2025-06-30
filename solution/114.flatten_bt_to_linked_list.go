package solution

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// recursively flatten left and right subtree
	Flatten(root.Left)
	Flatten(root.Right)

	// store left and right subtree
	left := root.Left
	right := root.Right

	// move subtree to the left
	root.Left = nil
	root.Right = left

	// find the tail of the new subtree
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// attach the original subtree
	current.Right = right
}
