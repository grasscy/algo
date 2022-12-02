package t1_199

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre, end := dummy, dummy
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil

		pre.Next = reverse(start)
		start.Next = next
		pre = start

		end = pre
	}
	return dummy.Next
}

func reverse(start *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
