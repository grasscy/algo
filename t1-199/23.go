package pass

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) > 2 {
		pre := lists[:2]
		p1 := mergeKLists(pre)
		return mergeKLists(append(lists[2:], p1))
	}
	dummy := &ListNode{}
	cur := dummy
	left := lists[0]
	right := lists[1]
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	for left != nil {
		cur.Next = left
		left = left.Next
		cur = cur.Next
	}
	for right != nil {
		cur.Next = right
		right = right.Next
		cur = cur.Next
	}
	return dummy.Next
}
