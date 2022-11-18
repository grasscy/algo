package pass

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		n := &ListNode{Val: list1.Val}
		if list1.Val > list2.Val {
			n.Val = list2.Val
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}
		cur.Next = n
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	// opt
	return head.Next
}
