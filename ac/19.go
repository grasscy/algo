package ac

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy
	for n != 0 {
		first = first.Next
		n--
	}

	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}
	if first == nil {
		second.Next = nil
	} else {
		second.Next = second.Next.Next
	}

	return dummy.Next
}
