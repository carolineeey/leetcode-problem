package solution

func Connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		var prev *Node // pointer to remember the previous node in the current level

		for i := 0; i < size; i++ {
			node := queue[0]  // take the first node from the queue
			queue = queue[1:] // remove it from the queue

			if prev != nil {
				prev.Next = node
			}
			prev = node

			// Enqueue the current node's children (if they exist) so we can process the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
