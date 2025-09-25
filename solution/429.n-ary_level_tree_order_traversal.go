package solution

func levelOrder(root *NodeChildren) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*NodeChildren{root}

	if len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		res = append(res, level)
	}
	return res
}
