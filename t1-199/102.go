package pass

func levelOrder(root *TreeNode) [][]int {
	r := [][]int{}
	if root == nil {
		return r
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		n := len(q)
		lv := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			lv = append(lv, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		r = append(r, lv)
	}
	return r
}
