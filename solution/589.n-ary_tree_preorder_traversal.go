package solution

type NodeChildren struct {
	Val      int
	Children []*NodeChildren
}

func Preorder(root *NodeChildren) []int {
	var res []int
	var dfs func(node *NodeChildren)
	dfs = func(node *NodeChildren) {
		if root == nil {
			return
		}

		res = append(res, node.Val) // visit root
		for _, child := range node.Children {
			dfs(child) // visit children
		}
	}

	dfs(root)
	return res
}
