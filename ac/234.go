package ac

//opt 复制链表值到数组列表中。
//使用双指针法判断是否为回文。
func isPalindrome(head *ListNode) bool {
	cur := head
	var pre *ListNode
	for cur != nil {
		n := &ListNode{cur.Val, nil}
		n.Next = pre
		pre = n
		cur = cur.Next
	}
	cur = head
	p := pre
	for cur != nil {
		if cur.Val != p.Val {
			return false
		}
		cur = cur.Next
		p = p.Next
	}
	return true
}
