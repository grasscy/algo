package ac

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			next := cur.Next.Next
			cur.Next = next
		} else {
			cur = cur.Next
		}
	}
	return head
}
