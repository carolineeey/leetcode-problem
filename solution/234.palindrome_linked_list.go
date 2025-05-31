package solution

func IsPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// find middle using slow and fast pointer
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// compare first and second half
	first, second := head, prev
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}

	return true
}
