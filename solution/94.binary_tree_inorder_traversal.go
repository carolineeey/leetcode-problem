package solution

func InorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left) //traverse left subtree
		result = append(result, node.Val)
		dfs(node.Right) //traverse right subtree
	}

	dfs(root)
	return result
}
