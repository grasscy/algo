package pass

//var all []*TreeNode

var a *TreeNode
var b *TreeNode
var pre *TreeNode

func recoverTree(root *TreeNode) {
	a, b = nil, nil
	tra(root)
	a.Val, b.Val = b.Val, a.Val
}

func tra(node *TreeNode) {
	if node == nil {
		return
	}
	tra(node.Left)
	if pre == nil {
		pre = node
	} else {
		if pre.Val > node.Val {
			b = node
			if a == nil {
				a = pre
			}
		}
		pre = node
	}
	tra(node.Right)
}
