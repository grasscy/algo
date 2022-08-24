package ac

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

//public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
//if (root == null) return null;
//// 如果p,q为根节点，则公共祖先为根节点
//if (root.val == p.val || root.val == q.val) return root;
//// 如果p,q在左子树，则公共祖先在左子树查找
//if (find(root.left, p) && find(root.left, q)) {
//return lowestCommonAncestor(root.left, p, q);
//}
//// 如果p,q在右子树，则公共祖先在右子树查找
//if (find(root.right, p) && find(root.right, q)) {
//return lowestCommonAncestor(root.right, p, q);
//}
//// 如果p,q分属两侧，则公共祖先为根节点
//return root;
//}
//
//private boolean find(TreeNode root, TreeNode c) {
//if (root == null) return false;
//if (root.val == c.val) {
//return true;
//}
//
//return find(root.left, c) || find(root.right, c);
//}
