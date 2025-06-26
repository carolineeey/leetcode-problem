package solution

func PathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int

	var dfs func(node *TreeNode, sum int, path []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}

		// add current node value to path
		path = append(path, node.Val)

		// check if it is a leaf and node value match to target
		if node.Left == nil && node.Right == nil && sum == node.Val {
			// make copy of current path if valid path found
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		dfs(node.Left, sum-node.Val, path)
		dfs(node.Right, sum-node.Val, path)
	}

	dfs(root, targetSum, []int{})
	return result
}
