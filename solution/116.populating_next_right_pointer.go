package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		// connect a node's left child to its right child
		root.Left.Next = root.Right
		// if current node has a Next, link Right child to the left child of the neighbor
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}

		Connect(root.Left)
		Connect(root.Right)
	}

	return root
}
