package solution

// preorder: root → left → right
// inorder: left → root → right
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Map to store value -> index from inorder
	inorderIndex := make(map[int]int)
	for i, val := range inorder {
		inorderIndex[val] = i
	}

	var build func(preStart, inStart, inEnd int) *TreeNode
	build = func(preStart, inStart, inEnd int) *TreeNode {
		// If preStart is out of bounds or the inorder range is invalid, there’s no node
		if preStart >= len(preorder) || inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inIndex := inorderIndex[rootVal]

		// Left subtree size = inIndex - inStart
		root.Left = build(preStart+1, inStart, inIndex-1)

		// Right subtree starts after the left subtree
		root.Right = build(preStart+(inIndex-inStart)+1, inIndex+1, inEnd)

		return root
	}

	return build(0, 0, len(inorder)-1)
}
