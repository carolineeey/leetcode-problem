package solution

func RecoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if prev != nil && prev.Val > node.Val { // detect violation (if nodes are out of order)
			if first == nil {
				first = prev // the first wrong node (prev)
			}
			second = node // the second wrong node (node)
		}
		prev = node // update node

		inorder(node.Right)
	}

	inorder(root)

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
