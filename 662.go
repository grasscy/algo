package algo

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*pair{{root, 1}}
	r := 1
	for len(queue) != 0 {
		n := len(queue)
		r = max(r, queue[len(queue)-1].index-queue[0].index+1)

		for i := 0; i < n; i++ {
			p := queue[0]
			queue = queue[1:]
			if p.node.Left != nil {
				queue = append(queue, &pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				queue = append(queue, &pair{p.node.Right, p.index*2 + 1})
			}
		}
	}
	return r
}
