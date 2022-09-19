package ac

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			l := root.Left
			for l.Right != nil {
				l = l.Right
			}
			root.Left = deleteNode(root.Left, l.Val)
			l.Right = root.Right
			l.Left = root.Left
			root = l
		}

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
