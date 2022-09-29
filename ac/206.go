package ac

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return h
}
