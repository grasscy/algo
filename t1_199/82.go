package t1_199

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{-111, head}

	cur := head
	pre := dummy
	for cur != nil {
		next := cur.Next
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		if next != cur.Next {
			pre.Next = next
		} else {
			pre = cur
		}
		cur = next
	}
	return dummy.Next
}
