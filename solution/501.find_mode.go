package solution

func FindMode(root *TreeNode) []int {
	var result []int
	var prev *TreeNode
	maxCount, count := 0, 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// inorder traversal: left, node, right
		inorder(node.Left)

		// count current value
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}

		// update maxCount and result list
		if count > maxCount {
			maxCount = count
			result = []int{node.Val}
		} else if count == maxCount {
			result = append(result, node.Val)
		}

		prev = node
		inorder(node.Right)
	}

	inorder(root)
	return result
}
