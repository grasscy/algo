package pass

func partition(head *ListNode, x int) *ListNode {
	bigger := &ListNode{0, nil}
	bh := bigger
	cur := head
	dummy := &ListNode{-1, head}
	pre := dummy
	for cur != nil {
		if cur.Val >= x {
			bh.Next = &ListNode{cur.Val, nil}
			bh = bh.Next

			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	pre.Next = bigger.Next
	return dummy.Next
}
