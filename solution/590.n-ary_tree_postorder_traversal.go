package solution

func Postorder(root *NodeChildren) []int {
	var res []int
	var dfs func(node *NodeChildren)
	dfs = func(node *NodeChildren) {
		if root == nil {
			return
		}

		for _, child := range node.Children {
			dfs(child) // visit children
		}
		res = append(res, node.Val) // visit root
	}

	dfs(root)
	return res
}
