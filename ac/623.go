package ac

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		n := &TreeNode{Val: val, Left: root}
		return n
	}
	q := []*TreeNode{root}
	for i := 1; i < depth-1; i++ {
		n := len(q)
		for j := 0; j < n; j++ {
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
	for _, v := range q {
		l := v.Left
		r := v.Right
		v.Left = &TreeNode{Val: val}
		v.Right = &TreeNode{Val: val}
		v.Left.Left = l
		v.Right.Right = r
	}
	return root
}
