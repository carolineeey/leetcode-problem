package solution

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// create a dummy node that acts as the start of the merged list
	dummy := &ListNode{}
	current := dummy

	// traverse through both lists until one of them is exhausted
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	// if one list is exhausted, append the remaining nodes of the other list
	if list1 == nil {
		current.Next = list2
	} else if list2 == nil {
		current.Next = list1
	}

	// return the merged list starting from dummy.Next (skipping the dummy node)
	return dummy.Next
}
