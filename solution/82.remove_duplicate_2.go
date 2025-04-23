package solution

func DeleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		// detect duplicate
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		if prev.Next != curr { // check if the duplicate found
			prev.Next = curr.Next // skip the duplicate
		} else {
			prev = prev.Next // no duplicate, move prev
		}

		curr = curr.Next // advance curr in both cases
	}

	return dummy.Next
}
