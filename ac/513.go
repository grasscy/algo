package ac

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{}
	q = append(q, root)
	ans := root.Val
	for len(q) != 0 {
		n := len(q)
		ans = q[0].Val
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
	}
	return ans
}
