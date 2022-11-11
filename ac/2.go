package ac

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	yu := 0
	ans := &ListNode{}
	rr := ans

	for l1 != nil && l2 != nil {
		v := (l1.Val + l2.Val + yu) % 10
		yu = (l1.Val + l2.Val + yu) / 10
		ans.Next = &ListNode{v, nil}
		ans = ans.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		v := (l1.Val + yu) % 10
		yu = (l1.Val + yu) / 10
		ans.Next = &ListNode{v, nil}
		ans = ans.Next
		l1 = l1.Next
	}
	for l2 != nil {
		v := (l2.Val + yu) % 10
		yu = (l2.Val + yu) / 10
		ans.Next = &ListNode{v, nil}
		ans = ans.Next
		l2 = l2.Next
	}
	if yu != 0 {
		ans.Next = &ListNode{yu, nil}
	}
	return rr.Next

}
