package pass

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur

		pre = next.Next
		cur = cur.Next
	}
	return dummy.Next
}
