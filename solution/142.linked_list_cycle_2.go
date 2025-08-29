package solution

func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
