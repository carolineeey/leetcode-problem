package solution

import "math"

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// recursively get the max gain from left and right subtree,
		// but discard negatives (treat them as 0)
		leftGain := max(0, maxGain(root.Left))
		rightGain := max(0, maxGain(root.Right))

		// new path that includes both children + current node
		newPath := root.Val + leftGain + rightGain

		// update global max
		if newPath > maxSum {
			maxSum = newPath
		}

		return root.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
