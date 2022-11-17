package ac

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
