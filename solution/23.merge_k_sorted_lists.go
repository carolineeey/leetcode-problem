package solution

//func MergeKLists(lists []*ListNode) *ListNode {
//	dummy := &ListNode{}
//	current := dummy
//
//	for i := 1; i < len(lists); i++ {
//		if lists[i-1] != nil && lists[i] != nil {
//			if lists[i-1].Val <= lists[i].Val {
//				current.Next = lists[i-1]
//				lists[i-1] = lists[i-1].Next
//			} else {
//				current.Next = lists[i]
//				lists[i] = lists[i].Next
//			}
//		}
//
//		if lists[i-1] == nil {
//			current.Next = lists[i]
//		} else if lists[i] == nil {
//			current.Next = lists[i-1]
//		}
//	}
//
//	return dummy.Next
//}

// Merge k sorted linked lists using iterative pairwise merging
func MergeKLists(lists []*ListNode) *ListNode {
	// if there are no lists, return nil
	if len(lists) == 0 {
		return nil
	}

	// keep merging the lists in pairs until one list remains
	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			// merge pairs of lists
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, MergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}

	return lists[0]
}
