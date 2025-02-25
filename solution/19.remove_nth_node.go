package solution

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// create a dummy node to handle edge cases like removing the head node
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// move fast pointer n+1 steps ahead to keep the gap
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// move both fast and slow pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove the nth node from end
	slow.Next = slow.Next.Next

	return dummy.Next
}
