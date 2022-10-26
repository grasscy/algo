package ac

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	if root.Val < low || root.Val > high {
		root = del(root, low, high)
	}
	return root
}

func del(root *TreeNode, low int, high int) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return nil
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		if root.Val < low {
			r := root.Right
			for r.Left != nil {
				r = r.Left
			}
			r.Left = root.Left
			return r.Right
		} else if root.Val > high {
			l := root.Left
			for l.Right != nil {
				l = l.Right
			}
			l.Right = root.Right
			return root.Left
		}
	}
	return root
}
