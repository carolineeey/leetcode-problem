package solution

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // save next node
		curr.Next = prev  // reverse pointer
		prev = curr       // move prev forward
		curr = next       // move curr forward
	}

	return prev
}
