package solution

import "strconv"

func BinaryTreePaths(root *TreeNode) []string {
	var result []string

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		// build current path string
		if len(path) > 0 {
			path += "->"
		}
		path += strconv.Itoa(node.Val)

		// if the node is a leaf then append
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
		}

		dfs(node.Left, path)
		dfs(node.Right, path)
	}

	dfs(root, "")
	return result
}
