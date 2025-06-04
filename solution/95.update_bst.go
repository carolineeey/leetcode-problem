package solution

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return buildTrees(1, n)
}

func buildTrees(start, end int) []*TreeNode {
	var trees []*TreeNode
	if start > end {
		return []*TreeNode{nil}
	}

	for i := start; i <= end; i++ {
		leftTrees := buildTrees(start, i-1) // all numbers < i go to the left subtree
		rightTrees := buildTrees(i+1, end)  // all numbers > i go to the right subtree

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				trees = append(trees, root)
			}
		}
	}

	return trees
}
