package solution

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length, tail := 1, head
	for tail.Next != nil { // get length
		tail = tail.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	curr := head
	for i := 0; i < length-k-1; i++ { // find new tail
		curr = curr.Next
	}

	// rearrange pointer
	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
}
