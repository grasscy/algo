package ac

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	q := []*pair{{root, 0}}
	ans := 1
	for len(q) != 0 {
		n := len(q)
		ans = max(ans, q[len(q)-1].index-q[0].index+1)

		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t.node.Left != nil {
				q = append(q, &pair{t.node.Left, 2*t.index + 1})
			}
			if t.node.Right != nil {
				q = append(q, &pair{t.node.Right, 2*t.index + 2})
			}
		}
	}
	return ans
}
