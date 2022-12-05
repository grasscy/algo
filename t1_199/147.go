package t1_199

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-111111, nil}
	tail := head
	cur := head
	for cur != nil {
		next := cur.Next
		tail.Next = nil

		cur.Next = nil
		cc := dummy
		for cc != nil && cc.Next != nil && cur.Val > cc.Next.Val {
			cc = cc.Next
		}

		if cc.Next == nil {
			tail = cur
		}
		cn := cc.Next
		cc.Next = cur
		cur.Next = cn

		cur = next
	}

	return dummy.Next
}
