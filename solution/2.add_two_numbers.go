package solution

import "fmt"

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val: 0,
	}
	pointer := result

	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry + 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		fmt.Println("sum l1", sum)
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		fmt.Println("sum l2", sum)

		carry = sum / 10
		sum = sum % 10
		fmt.Println("sum", sum, "carry", carry)

		pointer.Next = &ListNode{Val: sum}
		pointer = pointer.Next
	}

	if carry == 1 {
		pointer.Next = &ListNode{Val: carry}
	}

	return result.Next
}
