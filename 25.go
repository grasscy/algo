package algo

func reverseKGroup(head *ListNode, k int) *ListNode {
	end := head
	for i := 1; i < k && end != nil; i++ {
		end = end.Next
	}
	if end == nil {
		return head
	}

	node := reverse(head, end)
	head.Next = reverseKGroup(end.Next, k)

	return node
}

func reverse(start *ListNode, end *ListNode) *ListNode {
	if start == end {
		return start
	}
	var pre *ListNode
	cur := start
	for pre != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
