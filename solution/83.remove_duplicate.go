package solution

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next //skip duplicate node
		} else {
			curr = curr.Next
		}
	}

	return head
}
