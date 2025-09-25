package solution

func PreorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val) // visit root
		dfs(root.Left)              // traverse left subtree
		dfs(root.Right)             // traverse right subtree
	}

	dfs(root)
	return res
}
