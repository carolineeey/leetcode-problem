package solution

func BuildTreeTraversal(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// Map value -> index from inorder for O(1) lookup
	inorderIndex := make(map[int]int)
	for i, val := range inorder {
		inorderIndex[val] = i
	}

	var helper func(inStart, inEnd, postEnd int) *TreeNode
	helper = func(inStart, inEnd, postEnd int) *TreeNode {
		if inStart > inEnd || postEnd < 0 {
			return nil
		}

		// Last element in postorder is the root
		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}

		// Find root index in inorder
		idx := inorderIndex[rootVal]

		// Right subtree size
		rightSize := inEnd - idx

		// Recursively build left and right subtrees
		root.Left = helper(inStart, idx-1, postEnd-rightSize-1)
		root.Right = helper(idx+1, inEnd, postEnd-1)

		return root
	}

	return helper(0, len(inorder)-1, len(postorder)-1)
}
