package ac

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head2 := head.Next
	odd := head
	even := head.Next
	for even != nil && even.Next != nil {
		next := odd.Next
		odd.Next = next.Next
		odd = odd.Next

		even.Next = next
		even = even.Next
		even.Next = nil
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head2.Next

	return head
}
