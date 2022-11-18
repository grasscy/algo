package pass

func minDepth(root *TreeNode) int {
	ans := 1
	if root == nil {
		return 0
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return ans
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans++
	}
	return ans
}
