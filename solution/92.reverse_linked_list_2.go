package solution

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// make prev ends up pointing to the node right before the sublist you need to reverse
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// curr is the first node of the sublist (at position left)
	curr := prev.Next
	var next *ListNode
	for i := 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
