package solution

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSymmetric(p.Left, q.Right) && isSymmetric(p.Right, q.Left)
}
