package solution

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// if there is only one node, make it a leaf node
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// find middle node using slow and fast pointer
	prev := &ListNode{}
	slow, fast := head, head
	prev.Next = head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // fast moves twice as fast as slow
		prev = slow
		slow = slow.Next // slow will stop at the middle node, which becomes the root
	}

	// disconnect left half to avoid containing the full list
	prev.Next = nil

	// create tree node
	root := &TreeNode{Val: slow.Val}

	root.Left = SortedListToBST(head)       // left child is built from left half of list
	root.Right = SortedListToBST(slow.Next) // right child from the half after the middle

	return root
}
