package solution

func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += RangeSumBST(root.Left, low, high)
	}

	if root.Val > high {
		sum += RangeSumBST(root.Right, low, high)
	}

	return sum
}
