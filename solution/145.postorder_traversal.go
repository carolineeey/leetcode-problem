package solution

func PostorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)              // traverse left subtree
		dfs(root.Right)             // traverse right subtree
		res = append(res, root.Val) // visit root
	}

	dfs(root)
	return res
}
