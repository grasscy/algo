package ac

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}
	pre := cur
	cur = cur.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
