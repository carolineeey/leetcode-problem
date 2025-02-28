package solution

func SwapPairs(head *ListNode) *ListNode {
	// create a dummy node that points to the head to handle edge cases like an empty list
	dummy := &ListNode{}
	dummy.Next = head
	point := dummy

	for point.Next != nil && point.Next.Next != nil {
		// identify the two nodes to swap
		swap1 := point.Next
		swap2 := point.Next.Next

		// perform the swap
		swap1.Next = swap2.Next
		swap2.Next = swap1
		point.Next = swap2

		// move point to the last swapped node (swap1 after the swap)
		point = swap1
	}

	return dummy.Next
}
