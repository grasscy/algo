package t1_199

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	ts := []*TreeNode{}
	cur := head
	for cur != nil {
		ts = append(ts, &TreeNode{cur.Val, nil, nil})
		cur = cur.Next
	}
	return b(ts, 0, len(ts)-1)
}

func b(ts []*TreeNode, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := ts[mid]
	root.Left = b(ts, l, mid-1)
	root.Right = b(ts, mid+1, r)
	return root
}
