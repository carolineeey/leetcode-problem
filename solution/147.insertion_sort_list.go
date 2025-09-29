package solution

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: -1} // dummy head for sorted list
	curr := head                // iterate original list

	for curr != nil {
		prev := dummy
		next := curr.Next // store next pointer (before re-linking)

		// find place to insert current node
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}

		// insert between prev and prev.Next
		curr.Next = prev.Next
		prev.Next = curr

		// move to next node in original list
		curr = next
	}

	return dummy.Next
}
