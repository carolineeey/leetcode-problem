package solution

func HasCycle(head *ListNode) bool {
	slowPtr, fastPtr := head, head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if fastPtr == slowPtr {
			return true
		}
	}

	return false
}
