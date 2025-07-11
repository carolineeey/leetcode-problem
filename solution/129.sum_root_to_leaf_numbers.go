package solution

func SumNumbers(root *TreeNode) int {
	var sumNumbers func(root *TreeNode, curr int) int
	sumNumbers = func(root *TreeNode, curr int) int {
		if root == nil {
			return 0
		}

		// accumulate the number as you go down
		curr = curr*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return curr
		}

		return sumNumbers(root.Left, curr) + sumNumbers(root.Right, curr)
	}

	return sumNumbers(root, 0)
}
