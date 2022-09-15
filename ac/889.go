package ac

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}

	lroot := preorder[1]
	index := 0
	for ; index < len(postorder); index++ {
		if postorder[index] == lroot {
			break
		}
	}
	llen := index + 1
	root.Left = constructFromPrePost(preorder[1:llen+1], postorder[:llen])
	root.Right = constructFromPrePost(preorder[1+llen:], postorder[llen:len(postorder)-1])
	return root
}
