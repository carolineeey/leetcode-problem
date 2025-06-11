package solution

func LargestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		maxVal := queue[0].Val
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node.Val > maxVal {
				maxVal = node.Val
			}

			// add to next queue node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, maxVal)
		queue = queue[levelSize:]
	}

	return result
}
